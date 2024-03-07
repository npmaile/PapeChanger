//
//  utility.swift
//  Pape Changer
//
//  Created by Nathaniel Maile on 11/21/23.
//

import Foundation

func GetDirectoryCandidates() -> [String]{
    let parent = GetCurrentDirOfDirs()
    var ret:[String] = []
    do{
        ret = try FileManager.default.contentsOfDirectory(atPath:parent)
    }catch{
        print(error)
        ret = ["error"]
    }
    ret.sort()
    let realret = ret.filter({!$0.hasPrefix(".")})
    let actualrealret = realret.map({parent + "/" + $0})
    return actualrealret
}

func GetCurrentDirOfDirs() -> String{
    let task = Process()
    task.arguments = ["get", "-dirs"]
    let helper = Bundle.main.path(forAuxiliaryExecutable: "papechanger")
    task.executableURL = URL(fileURLWithPath: helper!)
    task.standardInput = nil
    
    let outpipe = Pipe()
    task.standardOutput = outpipe
    task.launch()
    let outdata = outpipe.fileHandleForReading.readDataToEndOfFile()
    
    var output : [String] = []
    if var string = String(data: outdata, encoding: .utf8) {
        string = string.trimmingCharacters(in: .newlines)
        output = string.components(separatedBy: "\n")
    }
    return output[0]
}

func ChangePapeDir(to:String) -> Void{
    let task = Process()
    task.arguments = ["cd","-direct", to]
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
