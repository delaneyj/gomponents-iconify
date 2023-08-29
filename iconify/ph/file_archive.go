package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileArchive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m213.66 82.34l-56-56A8 8 0 0 0 152 24H56a16 16 0 0 0-16 16v176a16 16 0 0 0 16 16h144a16 16 0 0 0 16-16V88a8 8 0 0 0-2.34-5.66ZM160 51.31L188.69 80H160ZM200 216h-88v-16h8a8 8 0 0 0 0-16h-8v-16h8a8 8 0 0 0 0-16h-8v-16h8a8 8 0 0 0 0-16h-8v-8a8 8 0 0 0-16 0v8h-8a8 8 0 0 0 0 16h8v16h-8a8 8 0 0 0 0 16h8v16h-8a8 8 0 0 0 0 16h8v16H56V40h88v48a8 8 0 0 0 8 8h48v120Z"/>`),
		g.Group(children),
	)
}