package main; 

import (
   "fmt"
   "os"
   "log"
   "strconv"
   "bufio"
   "strings"
)


func main() {
  file, err := os.Open("input.txt")
  if err != nil{
    log.Fatal(err)
  } 
  defer file.Close()

  one_counter := 0
  scanner := bufio.NewScanner(file)
  var array_one []string
  var array_zero []string
  var line_counter int 
  for scanner.Scan(){
   results := strings.Split(scanner.Text(), "")
      if(results[0] == "1"){
        array_one = append(array_one, scanner.Text()) 
        one_counter += 1
      } else {
        array_zero = append(array_zero, scanner.Text())
      }
   line_counter += 1
  }

  var arrayListDominant []string
  var arrayListLeast []string
  lengthOfArray := line_counter / 2 
  if one_counter < lengthOfArray {
   arrayListDominant = array_zero 
   arrayListLeast = array_one
  } else if one_counter == lengthOfArray{
   arrayListDominant = array_one
   arrayListLeast = array_zero
  } else {
   arrayListDominant = array_one
   arrayListLeast = array_zero
  }

  for i := 1; i<12; i++{
  if len(arrayListDominant) != 1{
    arrayListDominant = passOver(arrayListDominant,i, "dominant")
   }
  if len(arrayListLeast) != 1{
    arrayListLeast = passOver(arrayListLeast,i, "least")
   }
  }
  fmt.Println("arrayListDominant", arrayListDominant)
  fmt.Println("arrayListLeast", arrayListLeast)
  first_num,_ := strconv.ParseInt(strings.Join(arrayListDominant,""), 2, 64) 
  second_num,_ :=  strconv.ParseInt(strings.Join(arrayListLeast, ""), 2, 64)
  fmt.Println(first_num * second_num)
}

func passOver(arrayList []string, character_num int, returnType string) []string{
   fmt.Println("in passover, arrayList =", arrayList)
   var array_zero []string
   var array_one []string
   var one_counter int
   for _, number := range arrayList{
   results := strings.Split(number, "")
   fmt.Println("results",results)
   fmt.Println("results character_num",results[character_num])
   if(results[character_num] == "1"){
     array_one = append(array_one, number) 
     one_counter += 1
   } else {
     array_zero = append(array_zero, number)
   } 
  }
  fmt.Println("one_counter", one_counter)

  zero_counter := len(arrayList) - one_counter
  if one_counter < zero_counter{
   if returnType == "dominant"{
   return array_zero
   } else {
   return array_one
   }
  
  } else if one_counter == zero_counter{
   if returnType == "dominant"{
   return array_one
   } else {
   return array_zero
   }
  } else {
   if returnType == "dominant"{
   return array_one
   } else {
   return array_zero
   }
  }

}
