package main

import (
	// "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"

	"fmt"
	// "os"
)

func main() {
	// fmt.Println("vim-go")
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := cloudwatch.New(sess)

	resp, _ := svc.DescribeAlarms(nil)
	for _, alarm := range resp.MetricAlarms {
		// fmt.Println(*alarm.AlarmName)
		fmt.Printf(*alarm.AlarmName + "\t" + *alarm.StateValue + "\n")
	}
}
