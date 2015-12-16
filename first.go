//this is the first go program in centos
package main

func main(){
	var str = "this is the normal string"
	var slice =[]rune(str)
	for i:=0;i<len(slice)/2;i++{
		slice[i],slice[len(slice)-1-i]=slice[len(slice)-1-i],slice[i]
	}

	println(str)
	println(string(slice))
}
