package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderColorOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 24v-4h20v4H2Zm2-6v-3.75l9.05-9.05l3.75 3.75L7.75 18H4Zm2-2h.9L14 8.95L13.05 8L6 15.1v.9Zm11.925-8.15l-3.75-3.75l1.8-1.8q.275-.3.7-.287t.7.287l2.35 2.35q.275.275.275.688t-.275.712l-1.8 1.8ZM6 16Z"/>`),
		g.Group(children),
	)
}