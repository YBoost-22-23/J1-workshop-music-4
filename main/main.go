package main

import (
	"streaming"
)

func main() {
	streaming.Search()
}

/*
var (
	query      = flag.String("query", "Ynov", "Search term")
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
*/
