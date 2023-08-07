ifeq ($(OS),Windows_NT)
	Win-CC = $(CC)
else
	Win-CC = x86_64-w64-mingw32-gcc
endif

$(info $(SHELL))

mkdir: 
	mkdir -p build/bin

build: mkdir
	go build -o build/bin/papeChanger main.go

windows-build:
	go build -o build/bin/papeChanger.exe main.go

release-mac: mkdir
	go get ./...
	mkdir -p ./build/release/MacOS/
	cp -R ./build/package/Mac/PapeChanger.app ./build/release/MacOS/PapeChanger.app
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=1 go build -o build/release/MacOS/PapeChanger.app/Contents/MacOS/arm64_papeChanger main.go
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 go build -o build/release/MacOS/PapeChanger.app/Contents/MacOS/amd64_papeChanger main.go
	create-dmg \
		--app-drop-link 450 200 \
		--icon "PapeChanger.app" 150 200\
		--volname "PapeChanger Installer" \
		--hide-extension "PapeChanger.app" \
		--window-size 600 400 \
		--background "./assets/MacOS/installer_background.png" \
		./build/release/MacOS/PapeChanger.dmg \
		./build/release/MacOS/PapeChanger.app

release-win:
	go build -o build/release/Win/papeChanger.exe -ldflags -H=windowsgui main.go
	go run extra/wxsgenerator/generator.go build/package/Win/papeChanger.wxs.templ > ./build/release/Win/papeChanger.wxs
	cp ".\assets/icon.ico" ".\build\release\win"
	candle.exe ".\build\release\win\papeChanger.wxs" -ext WixUtilExtension -ext wixUIExtension -arch x64 -o build/release/win/papeChanger.wixobj
	light.exe ".\build\release\win\papeChanger.wixobj" -b ".\build\release\win;.\assets" -ext wixUIExtension  -ext WixUtilExtension -spdb

clean:
	rm -rf build/bin/*
	rm -rf build/release/*
