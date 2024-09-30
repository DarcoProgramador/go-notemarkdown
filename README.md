# Go NoteMarkdown

Go NoteMarkdown es una aplicación web para subir archivos Markdown y renderizarlos en HTML con resaltado de sintaxis.

## Requisitos

- Go 1.22.2 o superior

## Instalación

1. Clona el repositorio:

    ```sh
    git clone https://github.com/Darcoprogramador/go-notemarkdown.git
    cd go-notemarkdown
    ```

2. Instala las dependencias:

    ```sh
    go mod download
    ```

## Uso

1. Ejecuta la aplicación:

    ```sh
    go run main.go
    ```

2. Abre tu navegador y navega a `http://localhost:8000`.

## Estructura del Proyecto

```plaintext
go-notemarkdown/
├── go.mod
├── go.sum
├── handlers/
│   └── handlers.go
├── main.go
├── routes/
│   └── routes.go
├── storage/
│   └── storage.go
│   └── temp/
├── utils/
│   ├── markdown.go
│   └── markdown_test.go
└── views/
    └── index.html
```

## Endpoints
- Get /api/notes : Ve el listado de notas existente
- Post /api/notes : Sube un archivo Markdown
- Get /api/notes/:filename : Renderiza la nota a un html
- Get /api/notes/:filename/check : Checa la gramatica de la nota consultando a una API externa (Aun no implementado)

## Ejemplo de Uso
**Subir un archivo Markdown**
```sh
curl -F "file=@path/to/your/file.md" http://localhost:8000/api/notes
```
**Renderizar un Archivo MD**
Después de subir el archivo, puedes renderizarlo navegando a:
```plaintext
http://localhost:8000/api/notes/filename.md
```

****
[Reto de roadmap.sh](https://roadmap.sh/projects/markdown-note-taking-app)