# J Notificator

This is simple checker and notificator to the Telegram

# Config

Config file have to have the name "config.yaml".
J_notif will search for the config file in
```shell
/etc/j_notif/config.yaml
$HOME/.j_notif/config.yaml
./config.yaml
```
Config.yaml have to have writable permission (it used for deduplication)

# How to run

```shell
j_notif
```
or
```shell
j_notif >> /var/log/j_notif.log
```

# How to stop

```shell
kill -9 <j_notif_PID>
```