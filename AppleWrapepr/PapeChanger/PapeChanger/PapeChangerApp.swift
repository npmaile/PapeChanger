//
//  PapeChangerApp.swift
//  PapeChanger
//
//  Created by Nathaniel Maile on 8/13/23.
//

import SwiftUI
import Foundation
import HotKey

@main
struct PapeChangerApp: App {
    @State private var showPicker: Bool = false
    let changePapeHotKey = HotKey(key: .w, modifiers: [.control], keyDownHandler: {ChangePape()})
    var body: some Scene {
        MenuBarExtra("PapeChanger", systemImage: "square.on.square.fill") {
            Button("Change Wallpaper", action: ChangePape)
            Menu("Choose a Wallpaper Directory"){
                DirectoryChooser()
            }
            Button("Pick Wallpaper", action: selectPape)
            Divider()
            Button("Quit") { NSApplication.shared.terminate(nil) }
        }.menuBarExtraStyle(.menu)
    }
}
