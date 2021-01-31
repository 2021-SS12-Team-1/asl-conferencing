package botcommands

import (
   "log"
   "io"
   "fmt"
   "context"
	"godiscordspeechbot/bot"
	"godiscordspeechbot/commands/botcommands/commandUtils"
	"github.com/bwmarrin/discordgo"
   "google.golang.org/api/option"
   speech "cloud.google.com/go/speech/apiv1"
   speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

func Listen(b *bot.Bot, ctx *discordgo.MessageCreate, args []string) {
	v, ok := b.Session.VoiceConnections[ctx.GuildID]

	if ok != true {
		_ = b.Say(ctx, "You need to make me !join first", 5)
		return
	}

   fmt.Println("Connecting to Google Speech Recognition API...")

   ctx2 := context.Background()
	jsonPath := "/Users/sabra/go/src/godiscordspeechbot/commands/botcommands/credentials.json"

	// Inititalize the client that will be speaking to Speech-To-Text with the
	// credentials found in the credentials.json file
	client, err := speech.NewClient(ctx2, option.WithCredentialsFile(jsonPath))
	if err != nil {
		// We got an error when we tried to do that .. let's log the error and return
		log.Printf("speech.NewClient: %v", err)
		return
	}

   stream, err := client.StreamingRecognize(ctx2)
   if err != nil {
      log.Fatal(err)
   }

   // Send the initial configuration message.
   if err := stream.Send(&speechpb.StreamingRecognizeRequest{
      StreamingRequest: &speechpb.StreamingRecognizeRequest_StreamingConfig{
           StreamingConfig: &speechpb.StreamingRecognitionConfig{
               Config: &speechpb.RecognitionConfig{
                  AudioChannelCount: 2,
                  Encoding:        speechpb.RecognitionConfig_LINEAR16,
                  SampleRateHertz: 16000,
                  LanguageCode:    "en-US",
               },
               SingleUtterance: true,
           },
      },
   }); err != nil {
      log.Fatal(err)
   }

	recv := make(chan *discordgo.Packet, 2)
   //send := make(chan string, 2)
   go commandUtils.ReceiveAndConvertPCM(v, recv)
	go commandUtils.Start(recv, stream)

	//v.Speaking(true)
	//defer v.Speaking(false)


	for {
      resp, err := stream.Recv()
      if err == io.EOF {
         break
      }
      if err != nil {
         log.Printf("Cannot stream results: %v",   err)
      }
      if err := resp.Error; err != nil {
         // Workaround while the API doesn't give a more informative error.
         if err.Code == 3 || err.Code == 11 {
            log.Print("WARNING: Speech recognition request exceeded limit of 60 seconds.")
         }
         log.Fatalf("Could not recognize: %v", err)
      }
      for _, result := range resp.Results {
         fmt.Printf("Result: %+v\n", result)
         b.Say(ctx, result.Alternatives[0].Transcript)
      }
		if !ok {
			return
		}
	}
}
