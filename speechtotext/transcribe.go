package main

import (
   "io"
   "os"
   "fmt"
   "log"
   "context"
	"io/ioutil"

   "google.golang.org/api/option"
   speech "cloud.google.com/go/speech/apiv1"
   speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

func main() {
   ctx := context.Background()
   jsonPath := "credentials.json"

   client, err := speech.NewClient(ctx, option.WithCredentialsFile(jsonPath))
   if err != nil {
      log.Fatalf("speech.NewClient: %v", err)
   }

   f := "a-ok.wav"
   Send(os.Stdout, client, f)
}

func Send(w io.Writer, client *speech.Client, filename string) error {
   ctx := context.Background()
   data, err := ioutil.ReadFile(filename)
   if err != nil {
      log.Fatalf("ioutil.ReadFile(%s): %v", filename, err)
   }

   // Send the contents of the audio file with the encoding and
   // and sample rate information to be transcripted.
   req := &speechpb.LongRunningRecognizeRequest{
      Config: &speechpb.RecognitionConfig{
         Encoding:        speechpb.RecognitionConfig_LINEAR16,
         SampleRateHertz: 16000,
         LanguageCode:    "en-US",
     },
     Audio: &speechpb.RecognitionAudio{
        AudioSource: &speechpb.RecognitionAudio_Content{Content: data},
     },
   }

   op, err := client.LongRunningRecognize(ctx, req)
   if err != nil {
      return err
   }

   resp, err := op.Wait(ctx)
   if err != nil {
      return err
   }

   // Print the results.
   for _, result := range resp.Results {
      for _, alt := range result.Alternatives {
         fmt.Fprintf(w, "\"%v\" (confidence=%3f)\n", alt.Transcript, alt.Confidence)
      }
   }
   return nil
}
