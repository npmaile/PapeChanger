//
//  PapeSelectorView.swift
//  PapeChanger
//
//  Created by Nathaniel Maile on 8/14/23.
//

import SwiftUI
import UniformTypeIdentifiers

struct PapeSelectorView: View {
    @Environment(\.dismiss) private var dismiss
    var body: some View {
        VStack{
            Text("Please Pick a Wallpaper")
            HStack{
                Button("Select", action: selectPape)
            }
        }
    }
    
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
                task.arguments = ["--setup", filename!]
                task.standardInput = nil
                task.launch()
                task.waitUntilExit()
                dismiss()
            }
        }
    }
    
}


struct PapeSelectorView_Previews: PreviewProvider {
    static var previews: some View {
        PapeSelectorView()
    }
}