package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	// Channels to choose from
	channels := []string{
		"lofi hip hop radio",
		"sythwave radio",
		"peaceful piano radio",
		"dark ambient radio",
		"lofi sleep radio",
	}

	// everything always must be tokyo night or catpuccin
	tview.Styles.PrimitiveBackgroundColor = tcell.NewRGBColor(24, 25, 38)      // Background
	tview.Styles.ContrastBackgroundColor = tcell.NewRGBColor(32, 34, 46)       // Darker Background
	tview.Styles.MoreContrastBackgroundColor = tcell.NewRGBColor(40, 42, 54)   // Even Darker Background
	tview.Styles.BorderColor = tcell.NewRGBColor(103, 109, 138)                // Comment
	tview.Styles.TitleColor = tcell.NewRGBColor(198, 208, 245)                 // Foreground
	tview.Styles.GraphicsColor = tcell.NewRGBColor(103, 109, 138)              // Comment
	tview.Styles.PrimaryTextColor = tcell.NewRGBColor(198, 208, 245)           // Foreground
	tview.Styles.SecondaryTextColor = tcell.NewRGBColor(130, 137, 174)         // Light Comment
	tview.Styles.TertiaryTextColor = tcell.NewRGBColor(76, 80, 112)            // Dark Comment
	tview.Styles.InverseTextColor = tcell.NewRGBColor(24, 25, 38)              // Background
	tview.Styles.ContrastSecondaryTextColor = tcell.NewRGBColor(198, 208, 245) // Foreground

	app := tview.NewApplication()

	list := tview.NewList().
		ShowSecondaryText(false)

	for _, channel := range channels {
		// Add channels to the list
		list.AddItem(channel, "", 0, nil)
	}

	list.SetDoneFunc(func() {
		app.Stop()
	})

	list.SetSelectedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		url := getStreamURL(mainText)
		openStreamlink(url)
		app.Stop()
	})

	list.SetBorder(true).SetTitle("Select Channel to Open in Streamlink").SetTitleAlign(tview.AlignLeft)

	// Add key event handler
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyRune && event.Rune() == 'q' {
			app.Stop()
		}
		return event
	})

	if err := app.SetRoot(list, true).Run(); err != nil {
		log.Fatalf("Error running application: %v", err)
	}
}

// getStreamURL returns the stream URL based on the channel name
func getStreamURL(channel string) string {
	switch channel {
	case "lofi hip hop radio":
		return "https://www.youtube.com/watch?v=jfKfPfyJRdk"
	case "sythwave radio":
		return "https://www.youtube.com/watch?v=4xDzrJKXOOY"
	case "peaceful piano radio":
		return "https://www.youtube.com/watch?v=4oStw0r33so"
	case "dark ambient radio":
		return "https://www.youtube.com/watch?v=S_MOd40zlYU"
	case "lofi sleep radio":
		return "https://www.youtube.com/watch?v=rUxyKA_-grg"
	default:
		return ""
	}
}

// openStreamlink runs the streamlink command with the given URL
func openStreamlink(url string) {
	if url == "" {
		fmt.Println("No URL provided.")
		return
	}

	cmd := exec.Command("streamlink", url, "480p")
	if err := cmd.Start(); err != nil {
		log.Fatalf("Error starting streamlink: %v", err)
	}
	fmt.Printf("Opening stream: %s\n", url)
}
