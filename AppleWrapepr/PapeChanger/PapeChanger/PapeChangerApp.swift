//
//  PapeChangerApp.swift
//  PapeChanger
//
//  Created by Nathaniel Maile on 8/13/23.
//

import SwiftUI
import Foundation

@main
struct PapeChangerApp: App {
    @State private var showPicker: Bool = false
    @Environment(\.openWindow) private var openWindow
    @Environment(\.dismiss) private var dismiss
    var body: some Scene {
        MenuBarExtra("PapeChanger", systemImage: "\("p").circle") {
            Button("Change Wallpaper") { ChangePape() }
            Button("Change Wallpaper Directory") { ChangePapeDir() }
            Button("Pick Wallpaper") { openWindow(id:"selector") }
            Divider()
            Button("Quit") { NSApplication.shared.terminate(nil) }
        }
        Window("Wallpaper Selector", id: "selector"){
            PapeSelectorView()
        }
    }
}

func ChangePapeDir() -> Void{
    let task = Process()
    task.arguments = ["-c"]
    let helper = Bundle.main.path(forAuxiliaryExecutable: "papechanger")
    task.executableURL = URL(fileURLWithPath: helper!)
    task.standardInput = nil
    task.launch()
}

func ChangePape() -> Void{
    let task = Process()
    
    let helper = Bundle.main.path(forAuxiliaryExecutable: "papechanger")
    task.executableURL = URL(fileURLWithPath: helper!)
    task.standardInput = nil
    task.launch()
}

