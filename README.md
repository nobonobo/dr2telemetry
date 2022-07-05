# dr2telemetory

Telemetory Viewer for DiRT Rally 2.0

![](images/screenshot.png)

# Install

## STEP.1

edit: "%USERPROFILE%\Documents\My Games\DiRT Rally 2.0\hardwaresettings\hardware_settings_config.xml"

before:

```
<udp enabled="false" extradata="0" ip="127.0.0.1" port="20777" delay="1" />
```

after:

```
<udp enabled="true" extradata="3" ip="127.0.0.1" port="20777" delay="1" />
```

## STEP.2

Download and All Extract ZIP archive.

edit: params.json

```json
{
  "port": 20777,
  "lock2lock": 540,
  "window_x": 483,
  "window_y": 589,
  "window_w": 230,
  "window_h": 120
}
```

- The "port" must match between params.json and hardware_settings_config.xml.
- Set "lock2lock" to the operating range of your handle controller.
- other params saved automaticaly.

## Start

1. open `dr2telemetpry.exe`
2. Start and enjoy "DiRT Rally 2.0" !

# Behavier

- After 15 seconds, the telemetry display will be hidden behind.
- It will appear in the foreground when telemetry display is required.
