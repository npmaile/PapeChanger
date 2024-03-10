//
//  PapeChangerApp.swift
//  PapeChanger
//
//  Created by Nathaniel Maile on 8/13/23.
//

import SwiftUI
import Foundation
import KeyboardShortcuts

@main
struct PapeChangerApp: App {
    var body: some Scene {
        MenuBarExtra("PapeChanger", systemImage: "square.on.square.fill") {
            Button("Change Wallpaper", action: ChangePape)
            Menu("Choose a Wallpaper Directory"){
                DirectoryChooser()
            }
            Button("Pick Wallpaper", action: selectPape)
            Divider()
            SettingsLink()
            Divider()
            Button("Quit") { NSApplication.shared.terminate(nil) }
        }.menuBarExtraStyle(.menu)
        Settings{
                SettingsView()
        
        }
    }
}

@MainActor
final class AppState: ObservableObject {
    init() {
        KeyboardShortcuts.onKeyDown(for: .ChangePape, action: ChangePape)
    }
}

struct SettingsView: View{
    var body: some View {
        Form{
            KeyboardShortcuts.Recorder("Change Wallpaper:", name: .ChangePape)
        }
    }
}
