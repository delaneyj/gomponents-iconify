package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StraightenOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 18q-.825 0-1.413-.588T2 16V8q0-.825.588-1.413T4 6h16q.825 0 1.413.588T22 8v8q0 .825-.588 1.413T20 18H4Zm0-2h16V8h-3v4h-2V8h-2v4h-2V8H9v4H7V8H4v8Zm3-4h2h-2Zm4 0h2h-2Zm4 0h2h-2Zm-3 0Z"/>`),
		g.Group(children),
	)
}