package notification

import (
	"os/exec"

	"github.com/Alex222222222222/AppOcr/config"
	"github.com/Alex222222222222/AppOcr/files"
)

// I cannot find a pure go package to send macos system notification, seems object-c have package to deal with these things
// I decide to use terminal-notifier to do this for me.
// This method was pretty ugly and have many limitation.
// There was a golang repo called mac-driver, that may can send notification without using os.exec
// mac-driver does not support M1 macbook yet.

func Notify(c *config.UserConfig, title string, subtitle string, message string) error {
	files.SaveFile(files.IconPng, c.CacheDir+"/iconpng.png", false)
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
}
