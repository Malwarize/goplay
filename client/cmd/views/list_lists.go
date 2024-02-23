package views

import (
	"fmt"
	"net/rpc"

	"github.com/Malwarize/goplay/client/controller"
	"github.com/Malwarize/goplay/shared"
)

func PlayListsDisplay(client *rpc.Client) {
	playlists := controller.GetPlayListsNames(client)
	if len(playlists) == 0 {
		fmt.Println("No playlists")
		return
	}
	fmt.Println(GetTheme().ProgressStyle.Render("📼 Playlists:	"))
	for _, playlist := range playlists {
		fmt.Printf("\n   - %s\n", playlist)
	}
}

func PlayListSongsDisplay(name string, client *rpc.Client) {
	songs := controller.PlayListSongs(name, client)
	if len(songs) == 0 {
		fmt.Println("No songs in playlist")
		return
	}
	fmt.Println(GetTheme().ProgressStyle.Render("🎧 Playlist: ") + name)
	for index, song := range songs {
		fmt.Printf(
			"\n    %d : %s\n", index, shared.ViewParseName(song),
		)
	}
}
