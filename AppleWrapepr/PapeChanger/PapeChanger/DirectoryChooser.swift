//
//  DirectoryChooser.swift
//  Pape Changer
//
//  Created by Nathaniel Maile on 11/15/23.
//

import SwiftUI

struct DirectoryChooser: View {
    var body: some View {
        VStack{
            dirslist()
                .padding()
        }
    }
}

struct dirslist: View{
    var dirs: [Dir]
    
    init(){
        let listOfDirs = GetDirectoryCandidates()
        var list: [Dir] = []
        for d in listOfDirs{
            let nextDir = Dir(path:d)
            list.append(nextDir)
        }
        self.dirs = list
    }
    var body: some View{
        VStack{
            ForEach(dirs) {dir in
                dir
            }
        }
    }
}

struct Dir: View, Identifiable{
    var path: String
    var displayName: String
    var id: String { path }

    init(path: String){
        self.path = path
        let sp = path.split(separator: "/")
        let sub = sp.last ?? Substring()
        self.displayName = String(sub)
    }
    
    var body: some View{
        HStack{
            Button(action: changePapeDir){
                Text(displayName)
            }
        }
    }
    
    func changePapeDir(){
        ChangePapeDir(to:self.path)
    }
}
