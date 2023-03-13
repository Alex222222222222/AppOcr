# AppOcr

version: 1.0.1

A app for mac to capture screenshot, read qrcode, read barcode, read data matrix, or use ocr to scan text. 

## Configuration

Default configuration is at $HOME/.config/appocr.toml

```
# where to save your screen shot
ScreenShotSaveDir = "/Users/username/Downloads"

# default space ocr  
OCREngine = "Tesseract"

# space ocr api key
SpaceOCRAPI = "apikey"

# config file path, if you want to use your own path, please set path of your own config file in the default config file first
ConfigFilePath = "/Users/username/.config/appocr.toml"

# the executable path of tesseract, if you want to use tesseract, this is necessary
TesseractExecutablePath = "/opt/homebrew/bin/tesseract"

# where to save temporary file
CacheDir = "/Users/tesseract/Library/Caches/com.appocr"

# which language you want to use in the ocr language. Sea "Languages" in the readme.md
[[OcrLanguage]]                   
  SpaceOCR = "chs"
  Tesseract = "chi_sim"
```

## Install

terminal-notifier need to be installed first. And the executable path of terminal-notifier needed to be set in the config file to let the notification work fine.

## Permission

The "Screen Recording" permission in the macos system Preference in the "Security & Privacy" is required.

## Languages

Space OCR does not support multiple language, but you could set multiple OCR language by add entry in the OcrLanguage in the config file. 

```
[[OcrLanguage]]
  SpaceOCR = "chs"
  Tesseract = "chi_sim"

[[OcrLanguage]]
  SpaceOCR = "eng"
  Tesseract = "eng"
```

### SpaceOcr

- Arabic=ara
- Bulgarian=bul
- Chinese(Simplified)=chs
- Chinese(Traditional)=cht
- Croatian = hrv
- Czech = cze
- Danish = dan
- Dutch = dut
- English = eng
- Finnish = fin
- French = fre
- German = ger
- Greek = gre
- Hungarian = hun
- Korean = kor
- Italian = ita
- Japanese = jpn
- Polish = pol
- Portuguese = por
- Russian = rus
- Slovenian = slv
- Spanish = spa
- Swedish = swe
- Turkish = tur

> Language used for Space OCR. If no language is specified, Chinese "chs" is taken as default.

### Tesseract

- Afrikaans = afr
- Amharic = amh
- Arabic = ara
- Assamese = asm
- Azerbaijani = aze
- Azerbaijani - Cyrilic = aze_cyrl
- Belarusian = bel
- Bengali = ben
- Tibetan = bod
- Bosnian = bos
- Breton = bre
- Bulgarian = bul
- Catalan; Valencian = cat
- Cebuano = ceb
- Czech = ces
- Chinese simplified = chi_sim
- Chinese traditional = chi_tra
- Cherokee = chr
- Corsican = cos
- Welsh = cym
- Danish = dan
- German = deu
- Dhivehi = div
- Dzongkha = dzo
- Greek, Modern, 1453- = ell
- English = eng
- English, Middle, 1100-1500 = enm
- Esperanto = epo
- Math / equation detection module = equ
- Estonian = est
- Basque = eus
- Persian = fas
- Faroese = fao
- Filipino = fil
- Finnish = fin
- French = fra
- Frankish = frk
- French, Middle, ca.1400-1600 = frm
- West Frisian = fry
- Scottish Gaelic = gla
- Irish = gle
- Galician = glg
- Greek, Ancient, to 1453 = grc
- Gujarati = guj
- Haitian; Haitian Creole = hat
- Hebrew = heb
- Hindi = hin
- Croatian = hrv
- Hungarian = hun
- Armenian = hye
- Inuktitut = iku
- Indonesian = ind
- Icelandic = isl
- Italian = ita
- Italian - Old = ita_old
- Javanese = jav
- Japanese = jpn
- Kannada = kan
- Georgian = kat
- Georgian - Old = kat_old
- Kazakh = kaz
- Central Khmer = khm
- Kirghiz; Kyrgyz = kir
- Kurdish Kurmanji = kmr
- Korean = kor
- Korean vertical = kor_vert
- Lao = lao
- Latin = lat
- Latvian = lav
- Lithuanian = lit
- Luxembourgish = ltz
- Malayalam = mal
- Marathi = mar
- Macedonian = mkd
- Maltese = mlt
- Mongolian = mon
- Maori = mri
- Malay = msa
- Burmese = mya
- Nepali = nep
- Dutch; Flemish = nld
- Norwegian = nor
- Occitan post 1500 = oci
- Oriya = ori
- Orientation and script detection module = osd
- Panjabi; Punjabi = pan
- Polish = pol
- Portuguese = por
- Pushto; Pashto = pus
- Quechua = que
- Romanian; Moldavian; Moldovan = ron
- Russian = rus
- Sanskrit = san
- Sinhala; Sinhalese = sin
- Slovak = slk
- Slovenian = slv
- Sindhi = snd
- Spanish; Castilian = spa
- Spanish; Castilian - Old = spa_old
- Albanian = sqi
- Serbian = srp
- Serbian - Latin = srp_latn
- Sundanese = sun
- Swahili = swa
- Swedish = swe
- Syriac = syr
- Tamil = tam
- Tatar = tat
- Telugu = tel
- Tajik = tgk
- Thai = tha
- Tigrinya = tir
- Tonga = ton
- Turkish = tur
- Uighur; Uyghur = uig
- Ukrainian = ukr
- Urdu = urd
- Uzbek = uzb
- Uzbek - Cyrilic = uzb_cyrl
- Vietnamese = vie
- Yiddish = yid
- Yoruba = yor

> Language used for Tesseract. If no language is specified, Chinese "chs" is taken as default.