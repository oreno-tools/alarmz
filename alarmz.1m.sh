#!/usr/bin/env bash

export AWS_PROFILE=your-profile
export AWS_DEFAULT_REGION=ap-northeast-1
# for normal output mode
${HOME}/bin/alarmz
# for simple output mode
${HOME}/bin/alarmz -simple

# for tmux-powerline segment
# ${HOME}/bin/alarmz -output=/tmp/tmux-powerline-segments-cloudwatch_alarm_result.txt
# ${HOME}/bin/alarmz -simple -output=/tmp/tmux-powerline-segments-cloudwatch_alarm_result.txt
