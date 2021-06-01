package main


import (
	"fmt"
)

func main() {
N:=0
T:=[255]int{}
V:=0
F:=false
fmt.Println("Enter the length of your array:")
fmt.Scan(&N)

for n:=0;n<N;n++{
  G:=0
  fmt.Println("Enter elements of your array:")
  fmt.Scanln(&G)
  T[n]=G
}
fmt.Println("Enter the value for search:")
fmt.Scan(&V)
  for i:=0;i<len(T);i++{
    if T[i]==V {
      F=true
      fmt.Printf("Value %d is at %d index\n",V,i);
    }else{
      F=false
    }
  }
  if F== false{
    fmt.Printf("Value %d is not found in the array\n",V);
  } 

}