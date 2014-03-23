package main

import (
      "io/ioutil"
      "os"
      "flag"
      "fmt"
)

func visit(path string, f os.FileInfo, err error) error{
   fmt.Printf("%s\n", path)
   return nil
}

func main() {
   var showhidden = flag.Bool("a", false, "Shows hidden files")
   flag.Parse()
   root := flag.Arg(0)

   if root == "" {
      root = "."
   }
   filestuff, _ := ioutil.ReadDir(root)
   for _, onefile := range filestuff {
      if *showhidden || onefile.Name()[0] != '.' {
         fmt.Printf("%v ", onefile.Mode().String())
         fmt.Printf("%d ", onefile.Size())
         fileTime := onefile.ModTime().Format("Jan 02 15:04")
         fmt.Printf("%s ", fileTime)
         fmt.Printf("%s\n", onefile.Name())
      }
   }
}
