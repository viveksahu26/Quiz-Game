package main
import (
"flag"
"fmt"
"os"
"encoding/csv"
)

func main() {

  csvFile:=flag.String("csv", "problem.csv", "a problem file")
  
  flag.Parse()

  file,err:=os.Open(*csvFile)
  if err!=nil {
    exit(fmt.Sprintf("Failed to open CSV File:%s \n", *csvFile))
    
  }
  fmt.Println("csv:", *csvFile)
  r:=csv.NewReader(file)
  
  lines,err:=r.ReadAll()
  if err!=nil {
    exit("Failed to parse the CSV File.")
  }
  fmt.Println(lines)
}

func exit(message string){
  fmt.Print(message)
  os.Exit(10)
}