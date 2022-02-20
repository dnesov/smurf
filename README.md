# 🧙‍♂️ Smurf
A CLI tool for generating asset attributions file.

## 👷🏼‍♂️ Getting Started
As of today, there's no way to easily get Smurf from a package manager, so you need to build this repository manually with `go build .`.

## 🎉 Usage
Basic usage:
```
smurf -i "assets folder" [extensions]
```

### Example
If you want to create an `Attributions.md` for your image assets located in `Assets/` folder, go to your project root and run:
```
smurf -i Assets/ jpg png bmp
```

## 🤝 Contribution
This is my first Go project, so if you know how to improve it feel free to open pull requests.