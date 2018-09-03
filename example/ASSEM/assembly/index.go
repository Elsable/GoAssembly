package main

import (
	"github.com/AndrusGerman/GoAssembly"
)

//Como compilar -> Ejecuta en tu terminal el comando  "GOARCH=wasm GOOS=js go build"
//Como resultadon un archivo con el nombre assembli
func main() {
	//Se declaran las rutas
	GoAssembly.Ruta["index"] = RutaIndex()
	GoAssembly.Ruta["eventos"] = RutaEventos()
	GoAssembly.Ruta["variables"] = RutaVariables()
	GoAssembly.Ruta["extras"] = RutaExtra()
	//Run app
	GoAssembly.RunApp()
}

//TOda la pagina index
func RutaIndex() *GoAssembly.Page {
	var App GoAssembly.Page
	//Con la etiqueta ruta se agrega un id al elemento y el evento click para entrar a la ruta declarada entre la comillas
	App.Template = `
	<center>
		<h1>Pagina Inicio</h1>
		<button @rt="variables">Entra a Variables</button>
		<button @rt="eventos">Entrar a Eventos</button>
		<button @rt="extras">Extras</button>
		<button @rt="error">Not Found</button>
	</center>
	`
	return &App
}