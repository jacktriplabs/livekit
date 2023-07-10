<!--BEGIN_BANNER_IMAGE-->
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="/.github/banner_dark.png">
    <source media="(prefers-color-scheme: light)" srcset="/.github/banner_light.png">
    <img style="width:100%;" alt="The LiveKit icon, the name of the repository and some sample code in the background." src="/.github/banner_light.png">
  </picture>
  <!--END_BANNER_IMAGE-->

# LiveKit: Real-time video, audio and data for developers

[LiveKit](https://livekit.io) is an open source project that provides scalable, multi-user conferencing based on WebRTC.
It's designed to provide everything you need to build real-time video audio data capabilities in your applications.

LiveKit's server is written in Go, using the awesome [Pion WebRTC](https://github.com/pion/webrtc) implementation.

[![GitHub stars](https://img.shields.io/github/stars/livekit/livekit?style=social&label=Star&maxAge=2592000)](https://github.com/livekit/livekit/stargazers/)
[![Slack community](https://img.shields.io/endpoint?url=https%3A%2F%2Flivekit.io%2Fbadges%2Fslack)](https://livekit.io/join-slack)
[![Twitter Follow](https://img.shields.io/twitter/follow/livekitted)](https://twitter.com/livekitted)
[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/livekit/livekit)](https://github.com/livekit/livekit/releases/latest)
[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/livekit/livekit/buildtest.yaml?branch=master)](https://github.com/livekit/livekit/actions/workflows/buildtest.yaml)
[![License](https://img.shields.io/github/license/livekit/livekit)](https://github.com/livekit/livekit/blob/master/LICENSE)

## Features

- Scalable, distributed WebRTC SFU (Selective Forwarding Unit)
- Modern, full-featured client SDKs
- Built for production, supports JWT authentication
- Robust networking and connectivity, UDP/TCP/TURN
- Easy to deploy: single binary, Docker or Kubernetes
- Advanced features including:
    - [speaker detection](https://docs.livekit.io/guides/room/receive/#speaker-detection)
    - [simulcast](https://docs.livekit.io/guides/room/publish/#video-simulcast)
    - [end-to-end optimizations](https://blog.livekit.io/livekit-one-dot-zero/)
    - [selective subscription](https://docs.livekit.io/guides/room/receive/#selective-subscription)
    - [moderation APIs](https://docs.livekit.io/guides/server-api/)
    - [webhooks](https://docs.livekit.io/guides/webhooks/)
    - [distributed and multi-region](https://docs.livekit.io/deploy/distributed/)

## Documentation & Guides

https://docs.livekit.io

## Live Demos

- [LiveKit Meet](https://meet.livekit.io) ([source](https://github.com/livekit-examples/meet))
- [Spatial Audio](https://spatial-audio-demo.livekit.io/) ([source](https://github.com/livekit-examples/spatial-audio))
- Livestreaming from OBS Studio ([source](https://github.com/livekit-examples/livestream))
- [AI voice assistant using ChatGPT](https://livekit.io/kitt) ([source](https://github.com/livekit-examples/kitt))

## SDKs & Tools

### Client SDKs

Client SDKs enable your frontend to include interactive, multi-user experiences.

<table>
  <tr>
    <th>Language</th>
    <th>Repo</th>
    <th>
        <a href="https://docs.livekit.io/guides/room/events/#declarative-ui" target="_blank" rel="noopener noreferrer">Declarative UI</a>
    </th>
    <th>Links</th>
  </tr>
  <!-- BEGIN Template
  <tr>
    <td>Language</td>
    <td>
      <a href="" target="_blank" rel="noopener noreferrer"></a>
    </td>
    <td></td>
    <td></td>
  </tr>
  END -->
  <!-- JavaScript -->
  <tr>
    <td>JavaScript (TypeScript)</td>
    <td>
      <a href="https://github.com/livekit/client-sdk-js" target="_blank" rel="noopener noreferrer">client-sdk-js</a>
    </td>
    <td>
      <a href="https://github.com/livekit/livekit-react" target="_blank" rel="noopener noreferrer">React</a>
    </td>
    <td>
      <a href="https://docs.livekit.io/client-sdk-js/" target="_blank" rel="noopener noreferrer">docs</a>
      |
      <a href="https://github.com/livekit/client-sdk-js/tree/main/example" target="_blank" rel="noopener noreferrer">JS example</a>
      |
      <a href="https://github.com/livekit/client-sdk-js/tree/main/example" target="_blank" rel="noopener noreferrer">React example</a>
    </td>
  </tr>
  <!-- Swift -->
  <tr>
    <td>Swift (iOS / MacOS)</td>
    <td>
      <a href="https://github.com/livekit/client-sdk-swift" target="_blank" rel="noopener noreferrer">client-sdk-swift</a>
    </td>
    <td>Swift UI</td>
    <td>
      <a href="https://docs.livekit.io/client-sdk-swift/" target="_blank" rel="noopener noreferrer">docs</a>
      |
      <a href="https://github.com/livekit/client-example-swift" target="_blank" rel="noopener noreferrer">example</a>
    </td>
  </tr>
  <!-- Kotlin -->
  <tr>
    <td>Kotlin (Android)</td>
    <td>
      <a href="https://github.com/livekit/client-sdk-android" target="_blank" rel="noopener noreferrer">client-sdk-android</a>
    </td>
    <td>Compose</td>
    <td>
      <a href="https://docs.livekit.io/client-sdk-android/index.html" target="_blank" rel="noopener noreferrer">docs</a>
      |
      <a href="https://github.com/livekit/client-sdk-android/tree/main/sample-app/src/main/java/io/livekit/android/sample" target="_blank" rel="noopener noreferrer">example</a>
      |
      <a href="https://github.com/livekit/client-sdk-android/tree/main/sample-app-compose/src/main/java/io/livekit/android/composesample" target="_blank" rel="noopener noreferrer">Compose example</a>
    </td>
  </tr>
<!-- Flutter -->
  <tr>
    <td>Flutter (all platforms)</td>
    <td>
      <a href="https://github.com/livekit/client-sdk-flutter" target="_blank" rel="noopener noreferrer">client-sdk-flutter</a>
    </td>
    <td>native</td>
    <td>
      <a href="https://docs.livekit.io/client-sdk-flutter/" target="_blank" rel="noopener noreferrer">docs</a>
      |
      <a href="https://github.com/livekit/client-sdk-flutter/tree/main/example" target="_blank" rel="noopener noreferrer">example</a>
    </td>
  </tr>
  <!-- Unity -->
  <tr>
    <td>Unity WebGL</td>
    <td>
      <a href="https://github.com/livekit/client-sdk-unity-web" target="_blank" rel="noopener noreferrer">client-sdk-unity-web</a>
    </td>
    <td></td>
    <td>
      <a href="https://livekit.github.io/client-sdk-unity-web/" target="_blank" rel="noopener noreferrer">docs</a>
    </td>
  </tr>
  <!-- React Native -->
  <tr>
    <td>React Native (beta)</td>
    <td>
      <a href="https://github.com/livekit/client-sdk-react-native" target="_blank" rel="noopener noreferrer">client-sdk-react-native</a>
    </td>
    <td>native</td>
    <td></td>
  </tr>
  <!-- Rust -->
  <tr>
    <td>Rust</td>
    <td>
      <a href="https://github.com/livekit/client-sdk-rust" target="_blank" rel="noopener noreferrer">client-sdk-rust</a>
    </td>
    <td></td>
    <td></td>
  </tr>
</table>

### Server SDKs

Server SDKs enable your backend to generate [access tokens](https://docs.livekit.io/guides/access-tokens/),
call [server APIs](https://docs.livekit.io/guides/server-api/), and
receive [webhooks](https://docs.livekit.io/guides/webhooks/). In addition, the Go SDK includes client capabilities,
enabling you to build automations that behave like end-users.

| Language                | Repo                                                                                                | Docs                                                        |
|:------------------------|:----------------------------------------------------------------------------------------------------|:------------------------------------------------------------|
| Go                      | [server-sdk-go](https://github.com/livekit/server-sdk-go)                                           | [docs](https://pkg.go.dev/github.com/livekit/server-sdk-go) |
| JavaScript (TypeScript) | [server-sdk-js](https://github.com/livekit/server-sdk-js)                                           | [docs](https://docs.livekit.io/server-sdk-js/)              |
| Ruby                    | [server-sdk-ruby](https://github.com/livekit/server-sdk-ruby)                                       |                                                             |
| Java (Kotlin)           | [server-sdk-kotlin](https://github.com/livekit/server-sdk-kotlin)                                   |                                                             |
| Python (community)      | [tradablebits/livekit-server-sdk-python](https://github.com/tradablebits/livekit-server-sdk-python) |                                                             |
| PHP (community)         | [agence104/livekit-server-sdk-php](https://github.com/agence104/livekit-server-sdk-php)             |                                                             |

### Ecosystem & Tools

- [CLI](https://github.com/livekit/livekit-cli) - command line interface & load tester
- [Egress](https://github.com/livekit/egress) - export and record your rooms
- [Ingress](https://github.com/livekit/ingress) - ingest streams from RTMP / OBS Studio
- [Docker image](https://hub.docker.com/r/livekit/livekit-server)
- [Helm charts](https://github.com/livekit/livekit-helm)

## Install

We recommend installing [livekit-cli](https://github.com/livekit/livekit-cli) along with the server. It lets you access
server APIs, create tokens, and generate test traffic.

### MacOS

```shell
brew install livekit
```

### Linux

```shell
curl -sSL https://get.livekit.io | bash
```

### Windows

Download the [latest release here](https://github.com/livekit/livekit/releases/latest)

## Getting Started

### Starting LiveKit

Start LiveKit in development mode by running `livekit-server --dev`. It'll use a placeholder API key/secret pair.

```
API Key: devkey
API Secret: secret
```

To customize your setup for production, refer to our [deployment docs](https://docs.livekit.io/deploy/)

### Creating access token

A user connecting to a LiveKit room requires an [access token](https://docs.livekit.io/guides/access-tokens/). Access
tokens (JWT) encode the user's identity and the room permissions they've been granted. You can generate a token with our
CLI:

```shell
livekit-cli create-token \
    --api-key devkey --api-secret secret \
    --join --room my-first-room --identity user1 \
    --valid-for 24h
```

### Test with example app

Head over to our [example app](https://example.livekit.io) and enter a generated token to connect to your LiveKit
server. This app is built with our [React SDK](https://github.com/livekit/livekit-react).

Once connected, your video and audio are now being published to your new LiveKit instance!

### Simulating a test publisher

```shell
livekit-cli join-room \
    --url ws://localhost:7880 \
    --api-key devkey --api-secret secret \
    --room my-first-room --identity bot-user1 \
    --publish-demo
```

This command publishes a looped demo video to a room. Due to how the video clip was encoded (keyframes every 3s),
there's a slight delay before the browser has sufficient data to begin rendering frames. This is an artifact of the
simulation.

## Deployment

### Use LiveKit Cloud

LiveKit Cloud is the fastest and most reliable way to run LiveKit. Every project gets free monthly bandwidth and
transcoding credits.

Sign up for [LiveKit Cloud](https://cloud.livekit.io/).

### Self-host

Read our [deployment docs](https://docs.livekit.io/deploy/) for more information.

## Building from source

Pre-requisites:

- Go 1.18+ is installed
- GOPATH/bin is in your PATH

Then run

```shell
git clone https://github.com/livekit/livekit
cd livekit
./bootstrap.sh
mage
```

## Contributing

We welcome your contributions toward improving LiveKit! Please join us
[on Slack](http://livekit.io/join-slack) to discuss your ideas and/or PRs.

## License

LiveKit server is licensed under Apache License v2.0.

<!--BEGIN_REPO_NAV-->
<br/><table>
<thead><tr><th colspan="2">LiveKit Ecosystem</th></tr></thead>
<tbody>
<tr><td>Client SDKs</td><td><a href="https://github.com/livekit/components-js">Components</a> · <a href="https://github.com/livekit/client-sdk-js">JavaScript</a> · <a href="https://github.com/livekit/client-sdk-rust">Rust</a> · <a href="https://github.com/livekit/client-sdk-swift">iOS/macOS</a> · <a href="https://github.com/livekit/client-sdk-android">Android</a> · <a href="https://github.com/livekit/client-sdk-flutter">Flutter</a> · <a href="https://github.com/livekit/client-sdk-unity-web">Unity (web)</a> · <a href="https://github.com/livekit/client-sdk-python">Python</a> · <a href="https://github.com/livekit/client-sdk-react-native">React Native (beta)</a></td></tr><tr></tr>
<tr><td>Server SDKs</td><td><a href="https://github.com/livekit/server-sdk-js">Node.js</a> · <a href="https://github.com/livekit/server-sdk-go">Golang</a> · <a href="https://github.com/livekit/server-sdk-ruby">Ruby</a> · <a href="https://github.com/livekit/server-sdk-kotlin">Java/Kotlin</a> · <a href="https://github.com/agence104/livekit-server-sdk-php">PHP (community)</a> · <a href="https://github.com/tradablebits/livekit-server-sdk-python">Python (community)</a></td></tr><tr></tr>
<tr><td>Services</td><td><b>Livekit server</b> · <a href="https://github.com/livekit/egress">Egress</a> · <a href="https://github.com/livekit/ingress">Ingress</a></td></tr><tr></tr>
<tr><td>Resources</td><td><a href="https://docs.livekit.io">Docs</a> · <a href="https://github.com/livekit-examples">Example apps</a> · <a href="https://livekit.io/cloud">Cloud</a> · <a href="https://docs.livekit.io/oss/deployment">Self-hosting</a> · <a href="https://github.com/livekit/livekit-cli">CLI</a></td></tr>
</tbody>
</table>
<!--END_REPO_NAV-->
