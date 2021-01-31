package commandUtils

import (
   "io/ioutil"
   "os"
   "time"
   "log"
	"github.com/bwmarrin/discordgo"
   wave "github.com/zenwerk/go-wave"
   speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)


type SpeechMessage struct{
	PCM []int16
}

func Start(receive chan *discordgo.Packet, stream speechpb.Speech_StreamingRecognizeClient){
   audioFileName := "/Users/sabra/go/src/godiscordspeechbot/data/audio.wav"

   waveFile, _ := os.Create(audioFileName)

   inputChannels := 1
   sampleRate := 16000

   // setup Wave file writer
   param := wave.WriterParam{
      Out:           waveFile,
      Channel:       inputChannels,
      SampleRate:    sampleRate,
      BitsPerSample: 16,
   }

   waveWriter, _ := wave.NewWriter(param)

   for {
		select {
		case data := <- receive:
         _, err := waveWriter.WriteSample16(data.PCM)
         waveWriter.Close()
         time.Sleep(1 * time.Second)
         audio, err := ioutil.ReadFile(audioFileName)
         if err != nil {
            log.Fatalf("ioutil.ReadFile(%s): %v", audioFileName, err)
         }

         if err := stream.Send(&speechpb.StreamingRecognizeRequest{
            StreamingRequest: &speechpb.StreamingRecognizeRequest_AudioContent{
               AudioContent: audio,
            },
          }); err != nil {
             log.Printf("Could not send audio: %v", err)
         }
		}
	}

}
