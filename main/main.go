package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"google.golang.org/api/youtube/v3"

	"google.golang.org/api/googleapi/transport"
)

var (
	query      = flag.String("query", "Ynov lyon", "Search term")
	maxResults = flag.Int64("max-results", 100, "Max YouTube results")
)

const developerKey = "AIzaSyDsqSDIuvZC3PDglfoQkLQO8_As00il0D0"

func main() {
	flag.Parse()

	client := &http.Client{
		Transport: &transport.APIKey{Key: developerKey},
	}

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	// Make the API call to YouTube.
	call := service.Search.List([]string{"id", "snippet"}).
		Q(*query).
		MaxResults(*maxResults)
	response, err := call.Do()
	handleError(err, "")

	// Group video, channel, and playlist results in separate lists.
	videos := make(map[string]string)
	channels := make(map[string]string)
	playlists := make(map[string]string)

	// Iterate through each item and add it to the correct list.
	for _, item := range response.Items {
		switch item.Id.Kind {
		case "youtube#video":
			videos[item.Id.VideoId] = item.Snippet.Title + " | " + item.Snippet.ChannelTitle
		case "youtube#channel":
			channels[item.Id.ChannelId] = item.Snippet.Title
		case "youtube#playlist":
			playlists[item.Id.PlaylistId] = item.Snippet.Title
		}
	}
	printIDs("Videos", videos)
	printIDs("Channel", channels)
	printIDs("Playlist", playlists)

	//Demarrage du Serveur
	fmt.Println("localhost:8080")
	http.HandleFunc("/home", MainPage())
	http.ListenAndServe(":8080", nil)
}

// Print the ID and title of each result in a list as well as a name that
// identifies the list. For example, print the word section name "Videos"
// above a list of video search results, followed by the video ID and title
// of each matching video.
func printIDs(sectionName string, matches map[string]string) {
	fmt.Printf("%v:\n", sectionName)
	for _, title := range matches {
		fmt.Println(title)
	}
	fmt.Printf("\n\n")
}

func handleError(err error, message string) {
	if message == "" {
		message = "Error making API call"
	}
	if err != nil {
		log.Fatalf(message+": %v", err.Error())
	}
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./MainPage.html", "./index.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	tmpl.Execute(w, "")
}
