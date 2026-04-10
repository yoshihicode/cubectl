![cubectl demo: a rotating 3D cube in the terminal](screenshot.gif)
# 🧊 cubectl 
[![Release](https://img.shields.io/github/v/release/yoshihicode/cubectl)](https://github.com/yoshihicode/cubectl/releases/latest)
![Go Version](https://img.shields.io/github/go-mod/go-version/yoshihicode/cubectl)
![License](https://img.shields.io/github/license/yoshihicode/cubectl)

```bash
   ________  ______  ____________________
  / ____/ / / / __ )/ ____/ ____/__ __/ /
 / /   / / / / __  / __/ / /     / / / /
/ /___/ /_/ / /_/ / /___/ /___  / / / /___
\____/\____/_____/_____/\____/ /_/ /_____/
      _
    /_/|  Concept: "Control the Cube instead of Kubernetes cluster"
    |_|/
```
**"kubectl [kjuːb kəntróul] ... wait, CUBE control!?"**  
(Note: As a non-native English speaker, I took this pronunciation a bit too literally like `/kʊbɛ kəntróul/`. Hence, this tool was born.)

| Command | Pronunciation | Purpose |
| --- | --- | --- |
| `kubectl` | /kjuːb kəntróul/ | Controls Kubernetes clusters. |
| `cubectl` | /kjuːb kəntróul/ | Controls a cube instead of Kubernetes clusters. |

# ⚙️ Features
- 🧊 Renders a 3D cube instead of controlling Kubernetes clusters.
- 🔄 Rotate the cube with arrow keys or `wasd`.
- 🔍 Zoom in/out with `z` and `x`.
- 🚫 Absolutely no Kubernetes functionality included.

# 💾 Download
Download prebuilt binaries from the latest release:  
👉 [Get the latest binaries](https://github.com/yoshihicode/cubectl/releases/latest)

# 📦 Installation
## 🐧 Linux
```bash
wget https://github.com/yoshihicode/cubectl/releases/latest/download/cubectl_linux_amd64.tar.gz
tar -xzvf cubectl_linux_amd64.tar.gz
sudo mv cubectl /usr/local/bin/

# Run
cubectl
```
## 🍎🍺  macOS / Homebrew
```bash
brew tap yoshihicode/tap
brew install cubectl

# Run
cubectl
```
## 🪟 Windows
```powershell
Invoke-WebRequest -OutFile cubectl_windows_amd64.tar.gz https://github.com/yoshihicode/cubectl/releases/latest/download/cubectl_windows_amd64.tar.gz
tar -xzvf cubectl_windows_amd64.tar.gz

# Run
.\cubectl.exe
```

# 📘 Usage
```bash
cubectl controls cube instead of Kubernetes clusters.

Find more information at:
  https://github.com/yoshihicode/cubectl

Controls:
  Arrow keys or wasd: Rotate the cube
  z: Zoom in
  x: Zoom out
  Ctrl+C or Esc: Exit

Basic Cube Commands (Beginner):
  delete      Delete resources
  get         Display one or many cubes

Troubleshooting and Debugging Commands:
  describe    Describe resources
  logs        Print the logs for a cubectl



Flags:
  -h, --help   help for cubectl

Usage:
  cubectl [flags]

Usage:
  cubectl [command]

Use "cubectl <command> --help" for more information about a command.
```

# 🔌 Use as a kubectl plugin
You can integrate cubectl directly into your kubectl workflow with this simple wrapper:

```
echo -e '#!/bin/bash\ncubectl "$@"' > kubectl-cube
chmod +x kubectl-cube
sudo mv kubectl-cube /usr/local/bin/

# Now you're literally controlling a cube via kubectl!
kubectl cube 
kubectl cube logs -f
```

# 📚 References
## Inspiration for the idea
I was reading the following articles, and somehow ended up creating this joke command.  
- https://www.reddit.com/r/kubernetes/comments/7tzyla/poll_how_to_pronounce_kubectl/

## For 3D rendering
- https://codezine.jp/article/detail/38?p=3
- https://qiita.com/mochimkchiking/items/24fbbe93e0f7aa89edba
