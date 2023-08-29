package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 20h-2v-7H4v7H2V4h2v7h7V4h2v16Zm9-12v8h1.5v2H22v2h-2v-2h-5.5v-1.34l5-8.66H22Zm-2 3.133L17.19 16H20v-4.867Z"/>`),
		g.Group(children),
	)
}