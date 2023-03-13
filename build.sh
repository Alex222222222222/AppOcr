# a developer codeSigningCertificate is required to sign the app
codeSigningCertificate=""

rm -r ./AppOcr.app
rm -r ./AppOcr.tar.gz

mkdir -p ./AppOcr.app/Contents/MacOS/
mkdir -p ./AppOcr.app/Contents/Resources/

go mod tidy
go build main.go
cp main ./AppOcr.app/Contents/MacOS/AppOcr 
cp ./Info.plist ./AppOcr.app/Contents/
cp ./Icon.icns ./AppOcr.app/Contents/Resources/
codesign -s $codeSigningCertificate -f ./AppOcr.app/

tar -czvf ./AppOcr.tar.gz ./AppOcr.app