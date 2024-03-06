package main
import "fmt"

func main(){
	var smsSendingLimit int
	var costPerSms float64
	var hasPermission bool

	congrats := "Happy birthday 2 u"
	fmt.Printf(
		"%v %f %v %q\n",
		smsSendingLimit,
		costPerSms,
        hasPermission,
        congrats,
  
	)
}