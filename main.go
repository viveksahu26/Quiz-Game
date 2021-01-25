package main
import (
"flag"
"fmt"
"os"
"encoding/csv"
"strings"
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
  problems:=parseLines(lines)
  //fmt.Println(problems)
  correct := 0
  for i, p := range problems{
    fmt.Printf("Problem #%d: %s =  \n",i+1,p.q)
    var ans string
    fmt.Scanf("%s\n",&ans)
    if ans == p.a{
      correct++

    }
  }
  fmt.Printf("You scored %d out of %d", correct,len(problems))
}

func parseLines(lines [][]string) []problem{
  ret:=make([]problem,len(lines))
  for i, line := range lines{
    ret[i] = problem{
      q: line[0],
      a: strings.TrimSpace(line[1]),
    }
  }
  return ret
}

type problem struct{
  q string
  a string
}

func exit(message string){
  fmt.Print(message)
  os.Exit(10)
}