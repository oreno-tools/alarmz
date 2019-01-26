package main

import (
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"os"
	"strconv"
	"strings"
)

var (
	argVersion = flag.Bool("version", false, "バージョンを出力.")
	argTmux    = flag.Bool("tmux", false, "for tmux-powerline segment.")
)

const (
	appVersion = "0.0.1"
)

func main() {
	flag.Parse()

	if *argVersion {
		fmt.Println(appVersion)
		os.Exit(0)
	}

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
				fmt.Println(o)
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

	if *argTmux {
		fmt.Println("○: " + strconv.Itoa(len(oks)) + " !: " + strconv.Itoa(len(alarms)) + " ?: " + strconv.Itoa(len(insufficients)))
		os.Exit(0)
	} else {
		fmt.Printf("OK: " + strconv.Itoa(len(oks)) + " | color=green\n")
		fmt.Printf("ALARM: " + strconv.Itoa(len(alarms)) + " | color=red\n")
		fmt.Printf("INSUFFICIENT_DATA: " + strconv.Itoa(len(insufficients)) + " | color=purple\n")
		fmt.Println("---")
		fmt.Println(strings.Join(alarms, "\n"))
		fmt.Println("---")
		fmt.Println(strings.Join(insufficients, "\n"))
	}
}
