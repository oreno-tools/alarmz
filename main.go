package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"strconv"
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
	alarts := []string{}

	for {
		res, _ := svc.DescribeAlarms(params)
		for _, alarm := range res.MetricAlarms {
			// fmt.Println(alarm)
			if *alarm.StateValue == "OK" {
				oks = append(oks, "1")
			} else if *alarm.StateValue == "ALARM" {
				alarms = append(alarms, "1")
			} else if *alarm.StateValue == "INSUFFICIENT_DATA" {
				insufficients = append(insufficients, "1")
			}

			var alart string
			if *alarm.StateValue != "OK" {
				alart = *alarm.AlarmName + "\t" + *alarm.StateValue + "\n"
				alarts = append(alarts, alart)
			}
		}
		if res.NextToken == nil {
			break
		}
		params.SetNextToken(*res.NextToken)
		continue
	}

	fmt.Printf("OK: " + strconv.Itoa(len(oks)) + "\n")
	fmt.Printf("ALARM: " + strconv.Itoa(len(alarms)) + "\n")
	fmt.Printf("INSUFFICIENT_DATA: " + strconv.Itoa(len(insufficients)) + "\n")
	// fmt.Println(alarts)
}
