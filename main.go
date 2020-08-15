package main

import "github.com/fatih/color"
import "time"

func main(){
	color.Red("%s", time.Now())
}