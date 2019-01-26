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
	argVersion = flag.Bool("version", false, "Print version number.")
	argSimple  = flag.Bool("simple", false, "Simple ouput.")
	argOutput  = flag.String("output", "", "Set output file path.")
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

	icon_data := "| image=iVBORw0KGgoAAAANSUhEUgAAACQAAAAkCAYAAADhAJiYAAAAAXNSR0IArs4c6QAAAAlwSFlzAAAWJQAAFiUBSVIk8AAAAVlpVFh0WE1MOmNvbS5hZG9iZS54bXAAAAAAADx4OnhtcG1ldGEgeG1sbnM6eD0iYWRvYmU6bnM6bWV0YS8iIHg6eG1wdGs9IlhNUCBDb3JlIDUuNC4wIj4KICAgPHJkZjpSREYgeG1sbnM6cmRmPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4KICAgICAgPHJkZjpEZXNjcmlwdGlvbiByZGY6YWJvdXQ9IiIKICAgICAgICAgICAgeG1sbnM6dGlmZj0iaHR0cDovL25zLmFkb2JlLmNvbS90aWZmLzEuMC8iPgogICAgICAgICA8dGlmZjpPcmllbnRhdGlvbj4xPC90aWZmOk9yaWVudGF0aW9uPgogICAgICA8L3JkZjpEZXNjcmlwdGlvbj4KICAgPC9yZGY6UkRGPgo8L3g6eG1wbWV0YT4KTMInWQAAA89JREFUWAntl79rU1EUx/OzSZsWLUqgsXQqOGRwEBGnFtqpg04RB3Vw8A9w6FxwEwQd+w+4ZKxjoQV1LIgSsEuHIAhFoVrStPn1/Hxv3328l5e0SdpGBQ+cnHPuO/ec7zv3vHtvIpG/jKJDwKMccV+eFrr4j1CsS1Y/wIBL1wcBr8EMgVElJB/DT+Ab8A78C1ZuBx4K2VYYJ9sHWIktC8xtWHSRBTnO4P4mXPkaKSBVuO5K2V9gC9pKhi6G/AmUWAAarvTrN930gSp1azrX90KFwA2N7JK9JKMSa8lqrpRdgodKdtnGyPoOFgjLP9FvwaLAcmnATpR+3mQ/e8W9B9+Bv8KrsBpcVVRv9UwC2869TjYvmsvlxhYXF58uLy9P24lLS0v35+bmCrJXVlZCPWzX2vr7pUp8Jpqamqo7jvN8a2trdX5+/jvBkvClRCLxAlnc3Nz0V9HkOgmQcTiHnx+xWCwbjUavAs5pNpsKqT7qSAFAKiHcWlhY0BvcJcgeMWzjxVut1qONjY2S66cKamn8lWy3ldTMJ05DgIiZRISWSo6iAKBSqWTWngn5dDp9vVarRXg74yjJ203IcEttNzvz3P2x4Ewc/wP0GGBOPeUDgGwAJh40Go0IwI5g84ZUR76m3vRDa3R0NAXIsUqlYkFEMplMlLGjtbW1AxurX9mxdCopoBRLYMQJ19ZYRMtK9R4eHh5+i8fjJXhbUjYAXxmn41PeVXsXHQF1mO5VwfdsfGRkJIWdA2xW0rWvyAfQEn1Tr4BCPUEVmyyjEtbQpdRkA06b3sDUK6CTEihGjIYzsQAXAn/S5PZn5wGoPeaZ7P+ATitfrxXq9JWdFnug5x03xrZIvTXp8b7VNrV/s2OF+HRbfC2Kpp1ZR4TOIdke4cOQ5yO/pmzN9ZxQGDOxkF19/P7dKpRJJpMKliKB8ZfkLDPHiAbYc1Kcd54PvnHNYVy3RI+YN6FxjiJdPSJcPaSnPYc2JQAon887xWJRLh85Bq4hdU0wIAgcJ6n+UxnizCrj8xljD5ZPE/syPp/kYA9q1Pf1en2Wcd2rnWq1Ooncgf8NCjUsJ3mCk9xbGv9rlMtlh2UyjUM1ojMzM6H58t/d3fX6KJvNhvpUceRXKBQaOqilW7IBJY0Tl7O3lFf/wSssUyiYnTioJLZyNYg9CT9bX19/oyJwxzIX/kAPuUlmacJpXTWZ0HdeEnrzuuk6hLkZROgtczPY39/3EvVVASWw1E23zweVtkJeJr6eB3yWE1y4tK94yJXAX7Gz6MR1tA3A24rLvxKzXNL/Bgq89G+3jofPCxhIKgAAAABJRU5ErkJggg=="
	simple_output := "○ " + strconv.Itoa(len(oks)) + " x " + strconv.Itoa(len(alarms)) + " ? " + strconv.Itoa(len(insufficients))
	alarms_output := "---\n" + strings.Join(alarms, "\n") + "\n---\n" + strings.Join(insufficients, "\n")

	if *argSimple {
		fmt.Println(simple_output + icon_data)
		fmt.Println(alarms_output)
	} else {
		fmt.Println(icon_data)
		fmt.Println("---")
		fmt.Printf("OK: " + strconv.Itoa(len(oks)) + " | color=green\n")
		fmt.Printf("ALARM: " + strconv.Itoa(len(alarms)) + " | color=red\n")
		fmt.Printf("INSUFFICIENT_DATA: " + strconv.Itoa(len(insufficients)) + " | color=purple\n")
		fmt.Println(alarms_output)
	}

	if *argOutput != "" {
		file, err := os.OpenFile(*argOutput, os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()
		fmt.Fprintln(file, simple_output)
	}
}
