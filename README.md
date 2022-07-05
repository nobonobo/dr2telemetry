# dr2telemetory

Telemetory Viewer for DiRT Rally 2.0

![](images/screenshot.png)

#　 Prerequires

Runtime and Delvelepment

- [WebView2](https://developer.microsoft.com/en-us/microsoft-edge/webview2/)

Development

- [Go](https://go.dev/)
- [Wails v2](https://wails.io/)

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

Download [ZIP-file](https://github.com/nobonobo/dr2telemetory/releases/download/v1.0.0/dr2telemetory-win64-v1.0.0.zip) and All Extract ZIP archive.

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

# Build from Source

STEP.1 `go` install

```shell
winget install GoLang.Go
```

STEP.2 `wails` install

```shell
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

STEP.3 source clone and build

```shell
git clone https://github.com/nobonobo/dr2telemetory
cd dr2telemetory
wails build
```

Output: dr2telemetory/build/bin/dr2telemetory.exe

# Start

1. open `dr2telemetpry.exe`
2. Start and enjoy "DiRT Rally 2.0" !

# Behavier

- After 15 seconds, the telemetry display will be hidden behind.
- The location and size of the window is saved when the display is hidden.
- It will restore in the foreground when telemetry display is required.
