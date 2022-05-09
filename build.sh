version="0.1.0"
echo 'building for windows...'
GOOS=windows GOARCH=amd64 go build -o bin/jiggler-windows-$version .
echo 'building for linux...'
GOOS=linux GOARCH=amd64 go build -o bin/jiggler-linux-$version .
echo 'building for macos...'
GOOS=darwin GOARCH=amd64 go build -o bin/jiggler-darwin-$version .
