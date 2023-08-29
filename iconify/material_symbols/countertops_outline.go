package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CountertopsOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20v-8H2v-2h4q-.825 0-1.413-.588T4 8V4h6v4q0 .825-.588 1.413T8 10h8V7q0-.425-.288-.713T15 6q-.425 0-.713.288T14 7h-2q0-1.25.875-2.125T15 4q1.25 0 2.125.875T18 7v3h4v2h-2v8H4ZM6 8h2V6H6v2Zm0 10h5v-6H6v6Zm7 0h5v-6h-5v6ZM6 8h2h-2Zm6 7Z"/>`),
		g.Group(children),
	)
}