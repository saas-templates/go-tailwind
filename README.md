# Go + TailwindCSS

* Very simple Go application with HTTP server.
* Uses Go templates for pages (Located in `server/templates/`)
* Templates are styled using TailwindCSS.
  * Running `make tailwind` scans the template files and generates the optimised tailwind styles.
  * Output CSS file is written to `server/static/main.css`
* All files are embedded using `embed` package.

## Using

1. Clone the repository and rename.
2. Run `make` to build the single binary in `./bin`
