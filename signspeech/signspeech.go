// package signspeech
// this is a reworking of this python package:
// https://github.com/Cuperino/Signspeech
package signspeech

import (
	"context"
	"io/ioutil"
	"log"
	"strings"

	prose "github.com/jdkato/prose/v2"

	speech "cloud.google.com/go/speech/apiv1"
	"google.golang.org/api/option"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

// func Parse {{{
//
// Parses the string of transcribed audio
func Parse(text string) {
	// Create a new document with the default configuration:
	doc, _ := prose.NewDocument(text)
	sentences := doc.Sentences()

	tokenizer := prose.NewIterTokenizer()

	for _, sentence := range sentences {
		tokens := tokenizer.Tokenize(sentence)
		translation := Translate(tokens)

		var result string
		for _, word := range translation {
			result = result + word
		}

		Display(result, translation)
	}
} // }}}

// func Translate {{{
//
// Iterates through the parsed and tokenized transcription, adding only the necessary words
// to the the translation array that will be returned
func Translate(tokens []*prose.Token) []string {
	var translation []string
	var tone string
	tone = ""

	for _, token := range tokens {
		tag := token.Tag
		word := token.Text

		// Remove blacklisted words
		switch strings.ToLower(word) {
		case "is":
			continue
		case "be":
			continue
		case "the":
			continue
		case "of":
			continue
		case "are":
			continue
		case "by":
			continue
		case ",":
			continue
		case ";":
			continue
		case ":":
			continue
		}

		// Determine the tag for each word, add it to the translation array if it's a tag
		// we want to keep. See consts.go for more details on each part of speech tag
		switch tag {
		case CC:
			translation = append(translation, word)
		case CD:
			for _, letter := range word {
				translation = append(translation, letter)
			}
		case DT:
			continue
		case EX:
			continue // I'm not sure if this is the right way to handle this POS
		case FW:
			continue
		case IN:
			translation = append(translation, word)
		case JJ:
			translation = append(translation, word)
		case JJR:
			translation = append(translation, word)
		case JJS:
			translation = append(translation, word)
		case LS:
			for _, letter := range word {
				translation = append(translation, letter)
			}
		case MD:
			continue
		case NN:
			translation = append(translation, word)
		case NNS:
			translation = append(translation, word)
		case NNP:
			for _, letter := range word {
				translation = append(translation, letter)
			}
		case NNPS:
			for _, letter := range word {
				translation = append(translation, letter)
			}
		case PDT:
			continue
		case POS:
			translation = append(translation, word)
		case PRP:
			translation = append(translation, word)
		case PRPS:
			translation = append(translation, word)
		case RB:
			translation = append(translation, word)
		case RBR:
			translation = append(translation, word)
		case RBS:
			translation = append(translation, word)
		case RP:
			translation = append(translation, word)
		case SYM:
			if word == "?" {
				tone = "?"
			} else if word == "!" {
				tone = "!"
			}
			continue
		case TO:
			translation = append(translation, word)
		case UH:
			translation = append(translation, word)
		case VB:
			translation = append(translation, word)
		case VBD:
			translation = append(translation, word)
		case VBG:
			translation = append(translation, word)
		case VBN:
			translation = append(translation, word)
		case VBP:
			translation = append(translation, word)
		case VBZ:
			translation = append(translation, word)
		case WDT:
			continue
		case WP:
			translation = append(translation, word)
		case WPS:
			translation = append(translation, word)
		case WRB:
			translation = append(translation, word)
		default:
			continue
		}
	}

	return translation
} // }}}

// func Transcribe {{{
//
// Sends the specified audio file to Google Speech to Text to be transcribed
// Returns the translated text and speaker tag as well as any errors thrown
// during the process
func Transcribe(audiofile string) (string, int32, error) {
	var txt string
	var speaker int32

	ctx := context.Background()
	jsonPath := "credentials.json"

	// Inititalize the client that will be speaking to Speech-To-Text with the
	// credentials found in the credentials.json file
	client, err := speech.NewClient(ctx, option.WithCredentialsFile(jsonPath))
	if err != nil {
		// We got an error when we tried to do that .. let's log the error and return
		log.Printf("speech.NewClient: %v", err)
		return txt, speaker, err
	}

	// Let's read our audio data ..
	data, err := ioutil.ReadFile(audiofile)
	if err != nil {
		// We got an error when we tried to do that .. let's log the error and return
		log.Fatalf("ioutil.Readfile(%s): %v", audiofile, err)
		return txt, speaker, err
	}

	// Send the contents of the audio file with the encoding and sample rate information
	// to be transcribed
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

	// Create a Long Running Recognize operation
	op, err := client.LongRunningRecognize(ctx, req)
	if err != nil {
		// We got an error when we tried to do that .. let's log the error and return
		log.Printf("client.LongRunningRecognize: %v", err)
		return txt, speaker, err
	}

	// Let's wait for a response from the operation
	resp, err := op.Wait(ctx)
	if err != nil {
		// We got an error when we tried to do that .. let's log the error and return
		log.Printf("op.Wait: %v", err)
		return txt, speaker, err
	}

	// Let's use the result the recognizer gave the highest confidence
	result := resp.Results[0]
	highestalt := result.Alternatives[0]
	txt = highestalt.Transcript

	// Lets determine who said this phrase
	wordInfo := highestalt.Words[0]
	speaker = wordInfo.SpeakerTag

	// TO DO: Add code that will interact w/ the APIs to get speaker names??

	return txt, speaker, nil
} // }}}
