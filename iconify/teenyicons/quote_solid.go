package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuoteSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M2.5 1A1.5 1.5 0 0 0 1 2.5V10a4 4 0 0 0 4 4v-1a3 3 0 0 1-3-3V7h3.5A1.5 1.5 0 0 0 7 5.5v-3A1.5 1.5 0 0 0 5.5 1h-3Zm7 0A1.5 1.5 0 0 0 8 2.5V10a4 4 0 0 0 4 4v-1a3 3 0 0 1-3-3V7h3.5A1.5 1.5 0 0 0 14 5.5v-3A1.5 1.5 0 0 0 12.5 1h-3Z"/>`),
		g.Group(children),
	)
}