//repairDigikam - a program to repair my broken digikam database
//  A program to facilitate fixing my broken Digikam database.
//
//  Accepts two arguments:
//      -h --help  Help
//      -i --infile input file of bad filename data.  Default nullAlbum.dat
//      -o --outfile output file for corrected data.  Default namesCorrected.dat
//  
//
//  Copyright (c) 2017, Ben Bailey.  All rights reservied.
//  Use of this source code is governed by a BSD-stle
//  
//  License that can be found in the LICENSE file.
//
//  Additional insight is in repairDigikam.doc


package main

import (
    "flag"
	"bufio"
	"fmt"
	"os"
    "strings"
    "path/filepath"
    avl "github.com/emirpasic/gods/trees/avltree"
)

// My local def defaults.  Change for your use
// or override with command line values
var (
    defSearchDir = "/media/Photos/PictureAlbums"
    defInfile = "./nullAlbum.dat"
    defOutfile = "./namesCorrected.dat"
    RepairDigikam repairDigikam
)

// repairDigikam structure
// Central store of information in support of the application
type repairDigikam struct {

    Infile *string //name of file containing list of bad files obtained from sqlite3 search of digikam database
    Outfile *string //output file to use in in correcting digikam database.  key,filename
    SearchDir *string //Search this directory and obtain a list of all current files, store in TreeGoodFiles
    BadCounter int //number of bad filenames not found after searching the SearchDirectory
    FoundCounter int            //number of corrected (bad filenames corrected) found after searching SearchDirectory
    TreeGoodFiles *avl.Tree      //The binary search tree of filenames currently in the SearchDirectory
    CorrectedFileNames []string  //The list of bad filenames, after correcting " " to "_"

}


func check(e error) {
    if e != nil {
        panic(e)
    }
}

// loadDirectories
// Obtained insifht from walk.go golang example
// Walk and load the current directory structure from SearchDirectory
// Returns error
func loadDirectories()  error {

    fmt.Println()
    fmt.Println("-- Loading existing directories and files into memory in balanced binary tree structure.")

    //load the actual directory/fileNames as they exist now
    RepairDigikam.TreeGoodFiles = avl.NewWithStringComparator()
    e := filepath.Walk(*RepairDigikam.SearchDir, func(path string, f os.FileInfo, err error) error {
        _, file :=  filepath.Split(path)
        RepairDigikam.TreeGoodFiles.Put(file,path)
        return err
    })
    
    if e != nil {
        panic(e)
    }

    fmt.Println("\tFinished obtaining directories and good files.")
    fmt.Println("\tNumber of existing file names obtained: ", RepairDigikam.TreeGoodFiles.Size())
    fmt.Println()

    return nil 
}

// loadBadFiles loads the list of bad filenames from the disk
// For each line, it swaps spaces " " with underscore (" " -> "_")
// It then stores each corrected filename into an array, placing into struct RepairDigikam
// Nothing is returned.
func loadBadFiles()  { //from disk
    fmt.Println("-- Loading file containing bad names or directory locations that are causing problems with digikam database")
	badFileNamesHandle, err := os.Open(*RepairDigikam.Infile)
    check(err)

    badFileNamesScanner := bufio.NewScanner(badFileNamesHandle) //open a scanner for the file nullAlbum.dat
	defer badFileNamesHandle.Close()

    //Load list of bad file names that are in the digikam database and substitute " " with "_"
    RepairDigikam.CorrectedFileNames = make([]string, 0) //will contain the corrected file names

    tree := avl.NewWithStringComparator()

    for badFileNamesScanner.Scan() {
        badFileElement := badFileNamesScanner.Text() //list of file names that have spaces

        if(strings.Contains(badFileElement," ")) {
            badFileElement = strings.Replace(badFileElement," ","_",-1) //digikamCorrected inline
        }
        RepairDigikam.CorrectedFileNames = append(RepairDigikam.CorrectedFileNames, badFileElement) //append previously badFileName to goodFileNames
 
        digikamIndex := strings.Index(badFileElement,"|")
        digikamIndexVal := badFileElement[:digikamIndex]
        digikamIndexFileName := badFileElement[digikamIndex + 1:]
        tree.Put(digikamIndexVal,digikamIndexFileName)
    }
    fmt.Println("\tFinished loading file containing bad names or directory locations.")
    fmt.Println("\tNumber of bad names or locations: ", len(RepairDigikam.CorrectedFileNames))

    return 
}


//correctFileNames.  For each bad filename in RepairDigikam.CorrectedFileNames
//  -find in the binary tree
//  -wrtie to file
func correctFileNames()  {

    digikamCorrected, err := os.Create(*RepairDigikam.Outfile)
    defer digikamCorrected.Close()
    check(err)


    fmt.Println()
    fmt.Println("-- Unable to match following bad file names and locations:")

    for _, correctedFileName := range RepairDigikam.CorrectedFileNames {  //bad file names
        split := strings.Index(correctedFileName,"|")
        if(split <= -1) {
            continue
        }
        imageId := correctedFileName[:split]
        correctedFileName = correctedFileName[split+1:]

        
        value, foundit := RepairDigikam.TreeGoodFiles.Get(correctedFileName)
        if foundit {
            fileOutput:= fmt.Sprintf("%s,%s\n",imageId,value)
            digikamCorrected.WriteString(fileOutput)
            RepairDigikam.FoundCounter++
        } else {
            RepairDigikam.BadCounter++
            fmt.Println("\tDid not find it in tree! imageId: ", imageId, " correctedFileName: ", correctedFileName)
        }
	}
}

// parseCmdLine - provides options to this program, which allows defaults to be overwritten
// Not shown is '-h' but is available, because flag pkg provides default of '-h'
func parseCmdLine() {
    RepairDigikam.Infile    = flag.String("infile",  defInfile, "Provide an input file containing bad file names.")
    RepairDigikam.Outfile   = flag.String("outfile", defOutfile, "Provide an output file that will contain the corrected file names.")
    RepairDigikam.SearchDir = flag.String("searchDir", defSearchDir, "Program looks recursively for matchings beginning at this location.")

    flag.Parse()
}


func verifyFiles() {
    //does infile exist?  Exit on not exists
    _, err := os.Stat(*RepairDigikam.Infile)
    if err != nil {
        fmt.Println("Specified input file containing bad names not found.")
        fmt.Println("File ", *RepairDigikam.Infile, " not found!  err: ", err)
        os.Exit(1)
    } 

    //check if output file already exists to avoid unintended overwrite
    _, err = os.Stat(*RepairDigikam.Outfile)
    if err == nil { //true then file exists
        //file exists; overwrite?
        reader := bufio.NewReader(os.Stdin)
        for ;true; {
            fmt.Println()
            fmt.Print("Output file with the name ", *RepairDigikam.Outfile, " already exists.  Overrite? (y/n): ")
            answer, _ := reader.ReadString('\n')
            if !strings.Contains(strings.ToLower(answer), "y") {
                fmt.Println("Answer not y, exiting to prevent overwriting.")
                os.Exit(1)
            }
            break
        }
    }

    //check if searchDir directory exists.  Exit on not exists.
    _, err = os.Stat(*RepairDigikam.SearchDir)
    if err != nil { //directory location does not exist
        fmt.Println("Specified starting directory location ", *RepairDigikam.SearchDir, " not found.  Err: ", err)
        os.Exit(1)
    }
}

func main() {
    //get command line args
    parseCmdLine()

   
    fmt.Println()
    fmt.Println("Reading from input file: ", RepairDigikam.Infile)
    fmt.Println("Writing to output file: ", RepairDigikam.Outfile)
    fmt.Println("Searching beginning at directory: ", RepairDigikam.SearchDir)

    verifyFiles()

    fmt.Println()
    fmt.Println("Reading bad names from file: ", RepairDigikam.Infile)

    //do the work!
    loadBadFiles() 
    err := loadDirectories() //was goodFileList, err
    check(err)
    correctFileNames() 

    //show the results
    badNamesLen := len(RepairDigikam.CorrectedFileNames)
    correctedFiles := badNamesLen - RepairDigikam.BadCounter

    fmt.Println()
    fmt.Println("-- Starting to find new filename and location for bad files.")
    fmt.Println("\tWriting corrected data to ", *RepairDigikam.Outfile)

    fmt.Println()
    fmt.Println("-- Bad Files and Locations: ", badNamesLen, " Corrected Files and Locations: ", correctedFiles, " Files not corrected: ", RepairDigikam.BadCounter)
    fmt.Println("-- File Matched: ", RepairDigikam.FoundCounter)
    fmt.Println("Completed activity.")
    fmt.Println()
    fmt.Println("Job completed.")
    fmt.Println()
}
