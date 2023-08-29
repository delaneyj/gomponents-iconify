package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlopeDownhill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 19v3H2v-9l20 6M19.09 7.5l-.84 2.76l-10.12-3A2.89 2.89 0 0 0 5.1 4.5a2.89 2.89 0 0 0-2.76 3.03a2.897 2.897 0 0 0 5.23 1.58l10.12 3l-.84 2.78l4.82-2.6l-2.58-4.79Z"/>`),
		g.Group(children),
	)
}