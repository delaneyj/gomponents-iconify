package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GreaterThanEqual(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 5v2.156L22.531 14L6 20.844V23l20-8.219V13.22zm0 20v2h20v-2z"/>`),
		g.Group(children),
	)
}