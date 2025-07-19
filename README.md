# EDITORJS GO HEADER

`editorjs-go-header` is an official plugin for [`editorjs-go`](https://github.com/christiandenisi/editorjs-go) that provides rendering support for Editor.js `header` blocks into safe and semantic HTML output.

## ✨ Features

- ✅ Escapes HTML entities to prevent injection
- ✅ Renders heading levels from `<h1>` to `<h6>`
- 🧩 Plug-and-play integration with `editorjs-go`
- 🛡️ Validates header levels to avoid invalid HTML

---

## 📦 Installation

Install the plugin along with `editorjs-go`:

```bash
go get github.com/christiandenisi/editorjs-go
go get github.com/christiandenisi/editorjs-go-header
```

---

## 🚀 Usage

### Register the plugin in your main converter and convert

```go
package main

import (
    "fmt"
    "github.com/christiandenisi/editorjs-go"
    "github.com/christiandenisi/editorjs-go-plugin-header/header"
)

func main() {
    jsonData := []byte(`{
        "time": 1234567890,
        "version": "2.27.0",
        "blocks": [
            {
                "type": "header",
                "data": { "text": "This is a title", "level": 2 }
            }
        ]
    }`)

    conv := editorjs.New()
    editorjs.Register(conv, "header", header.RenderHeader)

    html, err := conv.Convert(jsonData)
    if err != nil {
        panic(err)
    }

    fmt.Println(html)
}
```

---

## 📌 Notes

- Header levels are validated: only levels 1 through 6 are allowed. Invalid values default to `<h2>`.
- Escaping is handled via `html.EscapeString(...)` to prevent raw HTML injection.

---

## 🧱 Compatibility

This plugin is compatible with:

- `editorjs-go` version `1.x`

---

## 👤 License

MIT – free to use, modify, and distribute.