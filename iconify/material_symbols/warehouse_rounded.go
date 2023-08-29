package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WarehouseRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 19V8.35q0-.625.338-1.125T3.25 6.5l8-3.2q.35-.15.75-.15t.75.15l8 3.2q.575.225.913.725T22 8.35V19q0 .825-.588 1.413T20 21h-4v-8H8v8H4q-.825 0-1.413-.588T2 19Zm7 2v-2h2v2H9Zm2-3v-2h2v2h-2Zm2 3v-2h2v2h-2Z"/>`),
		g.Group(children),
	)
}