//
//  PapeSelectorView.swift
//  PapeChanger
//
//  Created by Nathaniel Maile on 8/14/23.
//

import SwiftUI

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
        papePicker.begin{result in
            if result.rawValue == 0{
                return
            }
            if papePicker.url != nil && papePicker.url!.isFileURL{
                let filename = papePicker.url!.standardizedFileURL.absoluteString
                let task = Process()
                    task.launchPath = "papeChanger"
                    task.arguments = ["--setup", filename]
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
