package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"strconv"
	"strings"
)

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := cloudwatch.New(sess)

	params := &cloudwatch.DescribeAlarmsInput{}

	oks := []string{}
	alarms := []string{}
	insufficients := []string{}

	for {
		res, _ := svc.DescribeAlarms(params)
		for _, alarm := range res.MetricAlarms {
			if *alarm.StateValue == "OK" {
				o := *alarm.AlarmName + " | color=green"
				oks = append(oks, o)
			} else if *alarm.StateValue == "ALARM" {
				a := *alarm.AlarmName + " | color=red"
				alarms = append(alarms, a)
			} else if *alarm.StateValue == "INSUFFICIENT_DATA" {
				i := *alarm.AlarmName + " | color=purple"
				insufficients = append(insufficients, i)
			}

		}
		if res.NextToken == nil {
			break
		}
		params.SetNextToken(*res.NextToken)
		continue
	}

	fmt.Printf("OK: " + strconv.Itoa(len(oks)) + " | color=green\n")
	fmt.Printf("ALARM: " + strconv.Itoa(len(alarms)) + " | color=red\n")
	fmt.Printf("INSUFFICIENT_DATA: " + strconv.Itoa(len(insufficients)) + " | color=purple\n")
	fmt.Println("---")
	fmt.Println(strings.Join(alarms, "\n"))
	fmt.Println("---")
	fmt.Println(strings.Join(insufficients, "\n"))
}
