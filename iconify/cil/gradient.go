package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gradient(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M440 40H72a16 16 0 0 0-16 16v400a16 16 0 0 0 16 16h368a16 16 0 0 0 16-16V56a16 16 0 0 0-16-16Zm-16 192h-48v48h48v48h-48v40h-48v-40h-48v40h-48v-40h-48v40h-48v-40H88v-48h48v-48H88V72h336Z"/><path fill="currentColor" d="M136 280h48v48h-48zm48-48h48v48h-48zm48 48h48v48h-48zm48-48h48v48h-48zm48 48h48v48h-48zm-192-96h48v48h-48zm96 0h48v48h-48zm96 0h48v48h-48z"/>`),
		g.Group(children),
	)
}