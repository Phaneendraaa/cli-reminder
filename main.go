package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
)

func main() {
		if len(os.Args) < 3{
			fmt.Println("Please Enter  ==> <hh:mm> Reminder Note  <==")
			os.Exit(1);
		}
		
		parser := when.New(nil)
		parser.Add(en.All...)
		parser.Add(common.All...)

		result,err := parser.Parse(os.Args[1],time.Now())
		if err != nil {
			fmt.Println(err)
		}
		
		isFututreTime := time.Now().After(result.Time)
		if isFututreTime {
			fmt.Println("Reminder Time is in Past , please enter a future time")
			os.Exit(2)
		}

		diff := result.Time.Sub(time.Now())
		fmt.Println("Reminder will be displayed after",diff.Round(time.Second))
		time.Sleep(diff)



		err = beeep.Notify("Reminder",strings.Join(os.Args[2:]," "),"assets/information.png")


		if err!=nil{
			fmt.Println(err)
		}


		err = beeep.Beep(10,1000)
		if err!=nil{
			fmt.Println(err)
		}

		os.Exit(1)

}
