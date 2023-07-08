ifeq ($(OS),Windows_NT)
	Win-CC = $(CC)
else
	Win-CC = x86_64-w64-mingw32-gcc
endif

build: 
	go build -o build/bin/papeChanger main.go

windows-build:
	go build -o build/bin/papeChanger.exe main.go

release-mac:
	go get ./...
	GOOS=darwin GOARCH=arm64 go build -o build/bin/arm64_papeChanger main.go
	GOOS=darwin GOARCH=amd64 go build -o build/bin/amd64_papeChanger main.go
	cp arm64_papeChanger ./build/release/MacOS/PapeChanger.app/Contents/MacOS/arm64_papeChanger
	cp phonon_x86_64 ./build/release/MacOS/PapeChanger.app/Contents/MacOS/amd64_papeChanger
	create-dmg \
		--app-drop-link 100 300 \
		--icon "PapeChanger.app" 100 100\
		--volname "PapeChanger Installer" \
		--hide-extension "PapeChanger.app" \
		--window-size 1200 600 \
		--background "./assets/MacOS/background.png" \
		phonon.dmg \
		./release/MacOS/Phonon.app
release-win: windows-build
	go run extra/wxsgenerator/generator.go build/release/win/wix/papeChanger.wxs.templ > papeChanger.wxs
	candle.exe "papeChanger.wxs"  -ext WixUtilExtension -ext wixUIExtension -arch x64
	light.exe ".\papeChanger.wixobj" -b ".\release\win\wix" -ext wixUIExtension  -ext WixUtilExtension -spdb

checkout-submodules:
	git submodule update --init
