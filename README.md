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

Once you mistype `kubectl` as `cubectl`, you’ll end up controlling a cube instead of Kubernetes clusters.  
This tool makes your typos fun.  

# ⚙️ Features
- 🧊 Renders a 3D cube instead of controlling Kubernetes clusters.
- 🔄 Rotate the cube with arrow keys or `wasd`.
- 🔍 Zoom in/out with `z` and `x`.
- 🚫 Absolutely no Kubernetes functionality included.

# 💾 Download
👉 [Get the latest binaries](https://github.com/yoshihicode/cubectl/releases/latest)

# 🚀 Quick start
## 🐧 Linux
```bash
wget https://github.com/yoshihicode/cubectl/releases/latest/download/cubectl_linux_amd64.tar.gz
tar -xzvf cubectl_linux_amd64.tar.gz
sudo mv cubectl /usr/local/bin/
cubectl
```
## 🍎🍺  macOS / Homebrew
```bash
brew tap yoshihicode/tap
brew install cubectl
```
## 🪟 Windows
```powershell
Invoke-WebRequest -OutFile cubectl_windows_amd64.tar.gz https://github.com/yoshihicode/cubectl/releases/latest/download/cubectl_windows_amd64.tar.gz
tar -xzvf cubectl_windows_amd64.tar.gz
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

# 📚 References
## Inspiration for the idea
I was reading the following articles, and somehow ended up creating this joke command.  
- https://www.reddit.com/r/kubernetes/comments/7tzyla/poll_how_to_pronounce_kubectl/

## For 3D rendering
- https://codezine.jp/article/detail/38?p=3
- https://qiita.com/mochimkchiking/items/24fbbe93e0f7aa89edba
