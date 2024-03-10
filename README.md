# retro
play music with command line
![retro_music](https://github.com/Malwarize/retro/assets/130087473/c9824547-9b09-48fc-a113-e1a847793cca)

## 📦 Installation
$${\color{#AC3097}Install \space \color{white}Retro}$$ 
```sh
wget https://github.com/Malwarize/retro/releases/download/v0.0.27/install.sh
chmod +x install.sh
./install.sh 
```
$${\color{#AC3097}Uninstall \space \color{white}Retro}$$
```sh
# Todo
````


## 🕹️ Music Controls
$${\color{#AC3097}Play \space \color{white} Music}$$

```sh
retro play "Despacito - Luis Fonsi"
```

#### $${\color{#AC3097}Pause/Resume \space \color{white}Music}$$
```sh
retro pause  # ⏸️
retro resume # ▶ ️
```

#### $${\color{#AC3097}Next/Previous \space \color{white}Music Queue}$$
```sh
retro next     # ⏭️️
retro previous # ⏮️️
```

#### $${\color{#AC3097}Adjust \space Volume \space \color{white}Music Queue}$$
```sh
retro vol 50 # 🎚️ set volume to 50% 
retro vol 0  # 🔇 mute volume 
```

#### $${\color{#AC3097}Stop \space \color{white}Music Queue}$$
```sh
retro stop # 🛑
```


## 🚦️ Controls
#### $${\color{#AC3097}Logs \space \color{white}Control}$$
```sh
retro logs        # 📜 show all logs #last 200 lines 
retro logs info   # 📢 show all info logs 
retro logs error  # 🚫 show all error logs
retro logs warn   # ⚠️ show all warning logs
```

#### $${\color{#AC3097}Changing \space \color{white}Theme}$$
```sh
retro theme pink    #🧼 
retro theme purple  #🔮  
retro theme blue    #🌊
# TODO: retro theme custom 
```

#### $${\color{#AC3097}Command \space \color{white}Help}$$
```sh
retro help      #❓ show all commands
retro help play #❗ show play command help
```
## 📢 Acknowledgments
#### $${\color{#AC3097}retro \space \color{white}is \space  made  \space  by  \space  \color{#FF99EE} @Malwarize \color{white} \space with \space ❤️}$$ 

## ⚙️ Configuration 
#### $${\color{#AC3097}Config \space \color{white}File}$$
the config file is located by default in `~/.retro/config.json`
if not found, you can create it manually by 
```sh
mkdir -p ~/.retro
touch ~/.retro/config.json
```
$${\color{#AC3097}Default \space \color{white}Config}$$
```json

{
  "retro_path": "~/.retro/", // default path for main retro directory
  "path_ytldpl": "yt-dlp",   // path to yt-dlp
  "path_ffmpeg": "ffmpeg",   // path to ffmpeg
  "path_ffprobe": "ffprobe", // path to ffprobe
  "search_timeout": 60000000000, // 60 seconds timeout for search
  "theme": "pink",         // default theme 
  "db_path": "/home/xorbit/.retro/retro.db",
  "discord_rpc": true, // enable discord rpc (rich presence)
  "log_file": "~/.retro/retro.log" // log file path
}
```


you can change the config manually, easy to understand and modify.

$${\color{#AC3097}Note \space \color{white}that}$$

* ☝ ️ if you change the config file, its recommended to restart the retro service.
with `systemctl --user restart retro`
* ⚠️  the config file will override the default values.


## 📝 License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
