# alarmz

## これなに

* BitBar の CloudWatch Alarm プラグインです

## 使い方

plugin/alarmz.1m.sh を BitBar プラグインディレクトリに放り込んで下さい.

```sh
cp plugin/alarmz.1m.sh your-bit-bar-plugin-dir/
```

尚, 必要に応じて, 以下の環境変数の値を修正して下さい.

```sh
export AWS_PROFILE=your-profile
export AWS_DEFAULT_REGION=ap-northeast-1
```

また, API コールの間隔を調整したい場合には alarmz.1m.sh の `1m` を任意の値に変更して下さい. 例えば, 5 分間隔であれば `5m` と指定して下さい. 詳細については, BitBar の [README](https://github.com/matryer/bitbar#get-started) をご一読下さい.

## One more thing...

tmux-powerline に出力したい場合には, `-simple` オプションと `-ouput` オプション付きで実行するようにして下さい.

```sh
${HOME}/bin/alarmz -output=/tmp/tmux-powerline-segments-cloudwatch_alarm_result.txt
${HOME}/bin/alarmz -simple -output=/tmp/tmux-powerline-segments-cloudwatch_alarm_result.txt
```

tmux-powerline の segment では, `-ouput` オプションで指定したファイルを参照するようにすることで, 無駄な CloudWatch API コールを抑えることが出来ると思います.
