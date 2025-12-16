# MS Rewards – Bing Engine (Go)

A Go-based CLI tool that automates Bing searches to help compute Microsoft Rewards points. It provides a clean, colored terminal UI, supports debug mode, and can optionally open search URLs in your browser.

> ⚠️ **Disclaimer**
> This project is for educational purposes only. Automating searches may violate Microsoft Rewards terms of service. Use at your own risk.

---

## ✨ Features

* Colored, aligned terminal UI with box layouts
* Configurable number of searches
* Debug mode (no browser automation)
* Randomized or indexed search terms
* Clean log output for each query
* Cross-platform (Windows, Linux, macOS)

---

## 📦 Requirements

* Go **1.20+**
* A UTF-8 compatible terminal
* Default browser installed (for non-debug mode)

---

## 🚀 Installation

```bash
git clone https://github.com/yourusername/ms-rewards-bing-go.git
cd ms-rewards-bing-go
go build -o msrewards
```

---

## ▶️ Usage

```bash
./msrewards -num=10 -single=true -debug=false
```

### Available Flags

| Flag      | Description                        | Example        |
| --------- | ---------------------------------- | -------------- |
| `-num`    | Number of searches to perform      | `-num=10`      |
| `-single` | Use a single source file           | `-single=true` |
| `-debug`  | Print URLs without opening browser | `-debug=true`  |

---

## 🖥️ Sample Output

```text
┌───────────────────────────────────────────────────────┐
│ [X] MS Rewards - Bing Engine Golang Configuration [X] │
└───────────────────────────────────────────────────────┘
┌───────────────────────────────────────────────────────┐
│ [-] Search: 5                                         │
│ [-] Source: bing.txt                                  │
│ [-] Computing MS Rewards Points: 15                   │
└───────────────────────────────────────────────────────┘
┌───────────────────────────────────────────────────────┐
│ [-] Bing Queries                                      │
└───────────────────────────────────────────────────────┘
[0][invincible]  https://bing.com/search?q=invincible
[1][reduce]      https://bing.com/search?q=reduce
[2][nest]        https://bing.com/search?q=nest
[3][park]        https://bing.com/search?q=park
[4][label]       https://bing.com/search?q=label
```

---

## 🧠 How It Works

1. Reads search terms from a source file
2. Selects terms using an index list
3. Builds Bing search
