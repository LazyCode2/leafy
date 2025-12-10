
## 🛠️ Leafy: Minimal Static Site Generator

<img width="54" height="54" alt="image" src="https://github.com/user-attachments/assets/81833a3e-8755-46b4-aa5f-84eddc4065df" />

  **Leafy** is a minimal static site generator written in **Go**. It enables the rapid creation of a basic HTML site by processing Markdown content against an HTML template.

-----


## 📋 Requirements

  * **Go 1.20+** installed on your system.
  * Basic knowledge of **Markdown** and **HTML/Go Templates** (`.tmpl`).

-----

## 🚀 Installation

Install the latest version of the Leafy binary:

```bash
go install github.com/LazyCode2/leafy@latest
```

-----

## ⚙️ Usage Guide

### 1\. Initialize Project Structure

Run the following command in your desired project directory to create the necessary structure:

```bash
leafy --init
```

This command creates the following required directory structure:
```
  ├── content/ (For Markdown files)
  ├── output/ (Generated HTML site destination)
  └── template/ (For Go template files)
```

### 2\. Add Content

Place your Markdown source file in the `content` folder.
📂 Example Content Folder
```
content/
├── _index.md       # Homepage content
├── about.md        # Example page
└── contact.md      # Another example page
```

### 3\. Define Template

Create your HTML template file inside the `template` folder (e.g., `template/default.tmpl`).

The template **must** include the `{{ .Data }}` placeholder where the processed Markdown content will be inserted.

**Example Template** (`template/default.tmpl`):

```html
<!DOCTYPE html>
<html>
<head>
    <title>{{ .Title }}</title>
</head>
<body>
    <h1>My Static Site</h1>
    {{ .Data }} 
</body>
</html>
```

**Example Index Template** (`template/index.tmpl`):

```html
<h1>{{ .Title }}</h1>
<div>
  {{ range .Posts }}
    <p><a href="{{ .URL }}">{{ .Title }}</a></p>
  {{ end }}
</div>
```

### 4\. Build Site

Generate the final HTML file by executing the build command:

```bash
leafy --build
```

The resulting HTML file (`index.html`) will be placed in the **`output/`** directory.
