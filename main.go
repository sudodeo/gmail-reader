package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

var totalMessagesRead int

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func main() {
	ctx := context.Background()
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, gmail.GmailReadonlyScope, gmail.GmailModifyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := gmail.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Gmail client: %v", err)
	}

	user := "me"
	var listOfMessageIDs []string
	for {
		r, err := srv.Users.Messages.List(user).MaxResults(500).LabelIds("UNREAD").Do() //Users.Labels.List(user).Do()
		if err != nil {
			log.Fatalf("Unable to retrieve labels: %v", err)
		}
		messages := r.Messages
		if len(messages) == 0 {
			break
		}
		listOfMessageIDs = []string{}
		// fmt.Println(listOfMessageIDs)
		for _, msg := range messages {
			// srv.Users.Messages.Modify(user, msg.Id, &gmail.ModifyMessageRequest{})
			// if msg.Payload.Headers
			listOfMessageIDs = append(listOfMessageIDs, msg.Id)
		}
		// fmt.Println(listOfMessageIDs)
		err = srv.Users.Messages.BatchModify(user, &gmail.BatchModifyMessagesRequest{
			RemoveLabelIds: []string{"UNREAD"},
			Ids:            listOfMessageIDs,
		}).Do()
		if err != nil {
			log.Fatalf("Unable to retrieve Gmail client: %v", err)
		}
		totalMessagesRead += len(messages)
		fmt.Println("Read " + strconv.Itoa(len(messages)) + " messages\nTotal messages read: " + strconv.Itoa(totalMessagesRead) + "\n")
	}
}
