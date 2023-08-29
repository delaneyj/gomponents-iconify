package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clapperboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M20 3v14a1 1 0 0 1-1 1H1a1 1 0 0 1-1-1V3a1 1 0 0 1 1-1h1l3 3h2.5l-3-3h3l3 3H13l-3-3h3l3 3h2.5l-3-3H19a1 1 0 0 1 1 1z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}