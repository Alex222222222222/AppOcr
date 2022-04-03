package config

import (
	"errors"
	"os"

	"github.com/BurntSushi/toml"
)

type UserConfig struct {
	ScreenShotSaveDir              string // The directory the screenshot will be saved
	OCREngine                      string // The online ocr engine used to scan
	SpaceOCRAPI                    string // SpaceOcrAPI key, get from the space ocr website
	ConfigFilePath                 string
	OcrLanguage                    []*Language // the language that used in the ocr engine, default "chs" -> chinese simplified
	TesseractExecutablePath        string
	CacheDir                       string
	TerminalNotifierExecutablePath string
}

type Language struct {
	SpaceOCR  string
	Tesseract string
}

func ReadConfig(configFilePath string) (*UserConfig, error) {
	dc, err := GenerateDefaultConfig() // default config
	if err != nil {
		return nil, err
	}
	if configFilePath == "" {
		configFilePath = dc.ConfigFilePath
		return ReadConfig(configFilePath)
	}

	var conf UserConfig

	if _, err := os.Stat(configFilePath); errors.Is(err, os.ErrNotExist) {
		c, err := GenerateDefaultConfig()
		if err != nil {
			return nil, err
		}
		SaveConfig(c)
		return c, err
	}

	f, err := os.ReadFile(configFilePath)
	data := string(f)

	_, err = toml.Decode(data, &conf)
	if err != nil {
		return nil, err
	}

	if conf.ConfigFilePath != dc.ConfigFilePath {
		configFilePathNew := conf.ConfigFilePath

		f, err := os.ReadFile(conf.ConfigFilePath)
		data := string(f)

		_, err = toml.Decode(data, &conf)
		if err != nil {
			return nil, err
		}

		conf.ConfigFilePath = configFilePathNew
	}

	if conf.CacheDir == "" {
		conf.CacheDir = dc.CacheDir
	}

	if conf.ScreenShotSaveDir == "" {
		conf.ScreenShotSaveDir = dc.ScreenShotSaveDir
	}

	if conf.OCREngine == "" {
		conf.OCREngine = dc.OCREngine
	}

	if conf.OcrLanguage == nil {
		conf.OcrLanguage = dc.OcrLanguage
	}
	if conf.TerminalNotifierExecutablePath == "" {
		conf.TerminalNotifierExecutablePath = dc.TerminalNotifierExecutablePath
	}

	return &conf, nil

}

func SaveConfig(c *UserConfig) error {
	if c.ConfigFilePath == "" {
		dc, err := GenerateDefaultConfig()
		if err != nil {
			return err
		}

		c.ConfigFilePath = dc.ConfigFilePath
	}

	if _, err := os.Stat(c.ConfigFilePath); errors.Is(err, os.ErrNotExist) {
	} else {
		err := os.Remove(c.ConfigFilePath)
		if err != nil {
			return err
		}
	}

	f, err := os.Create(c.ConfigFilePath)
	if err != nil {
		return err
	}
	defer f.Close()

	err = toml.NewEncoder(f).Encode(*c)
	if err != nil {
		return err
	}

	return nil

}

func GenerateDefaultConfig() (*UserConfig, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	cachedir, err := os.UserCacheDir()
	if err != nil {
		return nil, err
	}

	c := UserConfig{
		ScreenShotSaveDir:              home + "/Downloads",
		OCREngine:                      "SpaceOCR",
		ConfigFilePath:                 home + "/.config/appocr.toml",
		CacheDir:                       cachedir + "/com.appocr",
		TerminalNotifierExecutablePath: "/opt/homebrew/bin/terminal-notifier",
		OcrLanguage: []*Language{
			{
				SpaceOCR:  "chs",
				Tesseract: "chi_sim",
			},
			{
				SpaceOCR:  "eng",
				Tesseract: "eng",
			},
		},
	}

	if _, err := os.Stat(c.CacheDir); errors.Is(err, os.ErrNotExist) {
		os.MkdirAll(c.CacheDir, 751)
	}

	return &c, nil
}
