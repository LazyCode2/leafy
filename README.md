# Leafy

Leafy is a minimal static site generator written in Go. It allows you to quickly generate a basic HTML site from Markdown content and a template.

---

## Requirements

- Go 1.20+ installed on your system
- Basic knowledge of Markdown and HTML + .tmpl 

---

## Install
```bash
go install github.com/LazyCode2/leafy@latest  
```
## How It Works

1. **Initialize a project structure**  
   Run the following command to create the folder structure:

   ```bash
   leafy --init
   ```

it will create 
*   ├── content
*   ├── output
*   └── template

2. Add your content
   Put Markdown files inside the content folder (content/content.md) !Only one for now.

3. Add a template
   Put a basic HTML template inside the template folder (e.g., template/default.tmpl) with a placeholder {{ .Data }} where the Markdown content will be inserted.

   **Example** : 
   ```html
    <!DOCTYPE html>
    <html>
    <head>
        <title>Leafy</title>
    </head>
    <body>
        {{ .Data }}
    </body>
    </html>
    ```

4. Build the site 
   Generate your HTML site with:
   ```bash
   leafy --build
   ```


