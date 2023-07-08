ifeq ($(OS),Windows_NT)
	Win-CC = $(CC)
else
	Win-CC = x86_64-w64-mingw32-gcc
endif

mkdir: 
	mkdir -p build/bin

build: mkdir
	go build -o build/bin/papeChanger main.go

windows-build: mkdir
	go build -o build/bin/papeChanger.exe main.go

release-mac: mkdir
	go get ./...
	mkdir -p ./build/release/MacOS/
	cp -R ./build/package/Mac/PapeChanger.app ./build/release/MacOS/PapeChanger.app
	GOOS=darwin GOARCH=arm64 go build -o build/release/PapeChanger.app/Contents/arm64_papeChanger main.go
	GOOS=darwin GOARCH=amd64 go build -o build/release/PapeChanger.app/Contents/amd64_papeChanger main.go
	create-dmg \
		--app-drop-link 100 300 \
		--icon "PapeChanger.app" 100 100\
		--volname "PapeChanger Installer" \
		--hide-extension "PapeChanger.app" \
		--window-size 1200 600 \
		--background "./assets/MacOS/installer_background.png" \
		phonon.dmg \
		./build/release/MacOS/PapeChanger.app
release-win: windows-build
	go run extra/wxsgenerator/generator.go build/release/win/wix/papeChanger.wxs.templ > papeChanger.wxs
	candle.exe "papeChanger.wxs"  -ext WixUtilExtension -ext wixUIExtension -arch x64
	light.exe ".\papeChanger.wixobj" -b ".\release\win\wix" -ext wixUIExtension  -ext WixUtilExtension -spdb

clean:
	rm -rf build/bin/*
