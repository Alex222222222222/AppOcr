package notification

import (
	"os/exec"

	"github.com/Alex222222222222/AppOcr/config"
	"github.com/Alex222222222222/AppOcr/files"
)

func Notify(c *config.UserConfig, title string, subtitle string, message string) error {
	files.SaveFile(files.IconPng, c.CacheDir+"/iconpng.png", false)
	//cmd := exec.Command(c.TerminalNotifierExecutablePath, "-title", title, "-subtitle", subtitle, "-message", message, "-sender", "com.appocr", "-appIcon", c.CacheDir+"/iconpng.png")
	cmd := exec.Command(c.TerminalNotifierExecutablePath, "-title", title, "-subtitle", subtitle, "-message", message, "-sender", "com.apple.screenshot.launcher", "-appIcon", c.CacheDir+"/iconpng.png")
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
