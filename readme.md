# HotReloader

[![goreleaser](https://github.com/danny270793/GoHotReloader/actions/workflows/release.yaml/badge.svg)](https://github.com/danny270793/GoHotReloader/actions/workflows/release.yaml)

Re-execute command when a file changes

## Installation

### From Github releases page

Go to [Release page](https://github.com/danny270793/gohotreloader/releases) then download the binary which fits your environment

### From terminal

Get the last version available on github

```bash
LAST_VERSION=$(curl https://api.github.com/repos/danny270793/GoHotReloader/releases/latest | grep tag_name | cut -d '"' -f 4)
```

Download the last version directly to the binaries folder

For Linux (linux):

```bash
curl -L https://github.com/danny270793/GoHotReloader/releases/download/${LAST_VERSION}/GoHotReloader_${LAST_VERSION}_linux_amd64.tar.gz -o ./hotreloader.tar.gz
```

Untar the downloaded file

```bash
tar -xvf ./hotreloader.tar.gz
```

Then copy the binary to the binaries folder

```bash
sudo cp ./GoHotReloader /usr/local/bin/hotreloader
```

Make it executable the binary

```bash
sudo chmod +x /usr/local/bin/hotreloader
```

```bash
hotreloader version
```

## Ussage

Run the binary and pass the binary and the file to re-execute every tima a file changes on the parent folder of the file

```bash
hotreloader ./example node ./index.js
```

## Follow me

[![YouTube](https://img.shields.io/badge/YouTube-%23FF0000.svg?style=for-the-badge&logo=YouTube&logoColor=white)](https://www.youtube.com/channel/UC5MAQWU2s2VESTXaUo-ysgg)
[![GitHub](https://img.shields.io/badge/github-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)](https://www.github.com/danny270793/)
[![LinkedIn](https://img.shields.io/badge/linkedin-%230077B5.svg?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/danny270793)

## LICENSE

[![GitHub License](https://img.shields.io/github/license/danny270793/GoHotReloader)](license.md)

## Version

![GitHub Tag](https://img.shields.io/github/v/tag/danny270793/GoHotReloader)
![GitHub Release](https://img.shields.io/github/v/release/danny270793/GoHotReloader)
![GitHub package.json version](https://img.shields.io/github/package-json/v/danny270793/GoHotReloader)

Last update 15/07/2024
