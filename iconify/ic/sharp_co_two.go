package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpCoTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 9h-5v6h5V9zm-1.5 4.5h-2v-3h2v3zM8 13v2H3V9h5v2H6.5v-.5h-2v3h2V13H8zm10.5 2.5v1h3V18H17v-3.5h3v-1h-3V12h4.5v3.5h-3z"/>`),
		g.Group(children),
	)
}