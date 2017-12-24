package main

import "fmt"

func main() {

	a:=[]int{5,3,1,2,4,6}

	for i:=len(a)-1;i>1;i-- {
		for j:=0;j<i;j++ {
			if(a[j]>a[j+1]){
				temp:=a[j]
				a[j]=a[j+1]
				a[j+1]=temp
			}
		}
	}

	fmt.Println(a)

}
