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
        MenuBarExtra("PapeChanger", systemImage: "square.on.square.fill") {
            Button("Change Wallpaper") { ChangePape() }
            Menu("Choose a Wallpaper Directory"){
                DirectoryChooser()
            }
            Button("Pick Wallpaper") { openWindow(id:"selector") }
            Divider()
            Button("Quit") { NSApplication.shared.terminate(nil) }
        }.menuBarExtraStyle(.menu)
        Window("Wallpaper Selector", id: "selector"){
            PapeSelectorView()
        }
    }
}

