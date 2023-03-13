package notification

import (
	"github.com/Alex222222222222/AppOcr/config"
	"github.com/Alex222222222222/AppOcr/files"

	gosxnotifier "github.com/deckarep/gosx-notifier"
)

// I cannot find a pure go package to send macos system notification, seems object-c have package to deal with these things
// I decide to use terminal-notifier to do this for me.
// This method was pretty ugly and have many limitation.
// There was a golang repo called mac-driver, that may can send notification without using os.exec
// mac-driver does not support M1 macbook yet.

func Notify(c *config.UserConfig, title string, subtitle string, message string) error {
	files.SaveFile(files.IconPng, c.CacheDir+"/iconpng.png", false)

	// At a minimum specifiy a message to display to end-user.
	note := gosxnotifier.NewNotification(message)

	// Optionally, set a title
	note.Title = title

	// Optionally, set a subtitle
	note.Subtitle = subtitle

	// Optionally, set a sound from a predefined set.
	note.Sound = gosxnotifier.Basso

	// Optionally, set a group which ensures only one notification is ever shown replacing previous notification of same group id.
	note.Group = "com.appocr"

	// Optionally, set a sender (Notification will now use the Safari icon)
	note.Sender = "com.apple.screenshot.launcher"

	// Optionally, specifiy a url or bundleid to open should the notification be
	// clicked.
	// note.Link = "http://www.yahoo.com" //or BundleID like: com.apple.Terminal

	// Optionally, an app icon (10.9+ ONLY)
	note.AppIcon = c.CacheDir + "/iconpng.png"

	// Optionally, a content image (10.9+ ONLY)
	// note.ContentImage = "gopher.png"

	//Then, push the notification
	err := note.Push()

	//If necessary, check error
	if err != nil {
		return err
	}

	return nil

	/*
		//cmd := exec.Command(c.TerminalNotifierExecutablePath, "-title", title, "-subtitle", subtitle, "-message", message, "-sender", "com.appocr", "-appIcon", c.CacheDir+"/iconpng.png")

		// the appIcon paramater of terminal-notifier just donnot work for unknown reason
		// I have tested to set the sender paramater to "com.appocr" but that also doesnot work
		// set the sender to "com.apple.screenshot.launcher" will display the system screenshot app icon
		cmd := exec.Command(c.TerminalNotifierExecutablePath, "-title", title, "-subtitle", subtitle, "-message", message, "-sender", "com.apple.screenshot.launcher", "-appIcon", c.CacheDir+"/iconpng.png")
		err := cmd.Run()
		if err != nil {
			return err
		}

		return nil
	*/
}
