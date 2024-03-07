//
//  PapeSelectorView.swift
//  PapeChanger
//
//  Created by Nathaniel Maile on 8/14/23.
//

import SwiftUI
import UniformTypeIdentifiers

func selectPape(){
    let papePicker = NSOpenPanel()
    papePicker.canChooseFiles = true
    papePicker.canChooseDirectories = false
    papePicker.allowsMultipleSelection = false
    papePicker.canDownloadUbiquitousContents = false
    papePicker.prompt = "Please choose wallpaper"
    papePicker.title = "Please Choose a Wallpaper"
    papePicker.allowedContentTypes = [UniformTypeIdentifiers.UTType.image]
    papePicker.begin{result in
        if result.rawValue == 0{
            return
        }
        if papePicker.url != nil && papePicker.url!.isFileURL{
            let filename = papePicker.url!.standardizedFileURL.path().removingPercentEncoding
            let task = Process()
            
            let helper = Bundle.main.path(forAuxiliaryExecutable: "papechanger")
            task.executableURL = URL(fileURLWithPath: helper!)
            task.arguments = [filename!]
            task.standardInput = nil
            task.launch()
            task.waitUntilExit()
        }
    }
}
