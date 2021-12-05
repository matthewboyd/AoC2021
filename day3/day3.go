package main; 

import (
   "fmt"
   "os"
   "log"
   "bufio"
   "strings"
   "strconv"
)


func main() {
  file, err := os.Open("input.txt")
  if err != nil{
    log.Fatal(err)
  } 
  defer file.Close()
  line_counter := 0
  scanner := bufio.NewScanner(file)
  var char_array [12]int
  for scanner.Scan(){
   results := strings.Split(scanner.Text(), "")
   for i := 0; i <= 11; i++{
   if(results[i] == "1"){
      char_array[i] += 1
   }
}
  line_counter += 1
  }
  fmt.Println("line_counter", line_counter)
  var gamma_string string 
  var epsilum string 
  for _, v := range char_array{
     if v < 500{
      gamma_string += "0"
      epsilum += "1"
      } else {
      gamma_string += "1"
      epsilum += "0"
     }
  }
  gamma_number, _ := strconv.ParseInt(gamma_string, 2,64)
  epsilum_number, _ := strconv.ParseInt(epsilum, 2,64)
  fmt.Println(gamma_number * epsilum_number)
}
