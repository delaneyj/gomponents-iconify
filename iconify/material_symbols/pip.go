package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 11V9h3.6L1.3 4.7l1.4-1.4L7 7.6V4h2v7H2Zm2 9q-.825 0-1.413-.588T2 18v-5h2v5h8v2H4Zm16-7V6h-9V4h9q.825 0 1.413.588T22 6v7h-2Zm-6 7v-5h8v5h-8Z"/>`),
		g.Group(children),
	)
}