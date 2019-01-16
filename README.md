# prometheus-webhook-dingtalk

Generating [DingTalk] notification from [Prometheus] [AlertManager] WebHooks.

## Building and running

### Build

```bash
make
```

其实make是编译不过的，里面用到了promu相关的信息，目前对promu工具不是特别了解（尝试过很多次，如果有人解决请教我下谢谢89424516@qq.com），因此建议按go src目录clone下来编译
go/src/github.com/timonwong/prometheus-webhook-dingtalk

### Running

```bash
./prometheus-webhook-dingtalk <flags>
```

## Usage

```
usage: prometheus-webhook-dingtalk --ding.profile=DING.PROFILE [<flags>]

Flags:
  -h, --help              Show context-sensitive help (also try --help-long and --help-man).
      --web.listen-address=":8060"
                          The address to listen on for web interface.
      --ding.profile=DING.PROFILE ...
                          Custom DingTalk profile (can be given multiple times, <profile>=<dingtalk-url>).
      --ding.timeout=5s   Timeout for invoking DingTalk webhook.
      --template.file=""  Customized template file (see template/default.tmpl for example)
      --log.level=info    Only log messages with the given severity or above. One of: [debug, info, warn, error]
      --version           Show application version.

```

### Build
扩展了目前的功能
0.3.1 
	1.支持at到用户的操作
	2.支持多用户at，调号分割
	3.tmpl支持基于字符串的正则函数，用于收敛预警消息到用户

