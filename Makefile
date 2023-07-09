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
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=1 go build -o build/release/MacOS/PapeChanger.app/Contents/arm64_papeChanger main.go
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 go build -o build/release/MacOS/PapeChanger.app/Contents/amd64_papeChanger main.go
	create-dmg \
		--app-drop-link 450 200 \
		--icon "PapeChanger.app" 150 200\
		--volname "PapeChanger Installer" \
		--hide-extension "PapeChanger.app" \
		--window-size 600 400 \
		--background "./assets/MacOS/installer_background.png" \
		./build/release/MacOS/phonon.dmg \
		./build/release/MacOS/PapeChanger.app
release-win: windows-build
	go run extra/wxsgenerator/generator.go build/release/win/wix/papeChanger.wxs.templ > papeChanger.wxs
	candle.exe "papeChanger.wxs"  -ext WixUtilExtension -ext wixUIExtension -arch x64
	light.exe ".\papeChanger.wixobj" -b ".\release\win\wix" -ext wixUIExtension  -ext WixUtilExtension -spdb

clean:
	rm -rf build/bin/*
	rm -rf build/release/*
