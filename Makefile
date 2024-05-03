ifeq ($(OS),Windows_NT)
	Win-CC = $(CC)
else
	Win-CC = x86_64-w64-mingw32-gcc
endif
SHELL = /bin/sh -l

$(info $(SHELL))

mkdir: 
	mkdir -p build/bin
	mkdir -p build/release/MacOS

build: mkdir
	go build -o build/bin/papeChanger main.go

windows-build:
	go build -o build/bin/papeChanger.exe main.go

build-mac: mkdir clean
	mkdir -p ./build/bin/MacOS/
	go get ./...
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=1 go build -o build/bin/MacOS/arm64_papeChanger main.go
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 go build -o build/bin/MacOS/amd64_papeChanger main.go
	lipo build/bin/MacOS/amd64_papeChanger -create build/bin/MacOS/arm64_papeChanger -output ./build/bin/MacOs/papechanger

archive-mac: mkdir clean
	xcodebuild \
	-workspace AppleWrapepr/PapeChanger/Pape\ Changer.xcodeproj/project.xcworkspace \
	-scheme PapeChanger \
	-configuration Release \
	-archivePath ./PapeChanger.xcarchive \
	archive

	mv ./PapeChanger.xcarchive/Products/Applications/Pape\ Changer.app ./build/release/MacOS/Pape\ Changer.app

create-dmg:
	create-dmg \
		--app-drop-link 450 200 \
		--icon "Pape Changer.app" 150 200\
		--volname "PapeChanger Installer" \
		--hide-extension "Pape Changer.app" \
		--window-size 600 400 \
		--background "./assets/MacOS/installer_background.png" \
		./build/release/MacOS/PapeChanger.dmg \
		./build/release/MacOS/Pape\ Changer.app
	
release-win:
	go build -o build/release/Win/papeChanger.exe -ldflags -H=windowsgui main.go
	go run extra/wxsgenerator/generator.go build/package/Win/papeChanger.wxs.templ > ./build/release/Win/papeChanger.wxs
	cp ".\assets/icon.ico" ".\build\release\win"
	candle.exe ".\build\release\win\papeChanger.wxs" -ext WixUtilExtension -ext wixUIExtension -arch x64 -o build/release/win/papeChanger.wixobj
	light.exe ".\build\release\win\papeChanger.wixobj" -b ".\build\release\win;.\assets" -ext wixUIExtension  -ext WixUtilExtension -spdb

clean:
	rm -rf build/bin/*
	rm -rf build/release/*
