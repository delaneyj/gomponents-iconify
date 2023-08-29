package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Windy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M21 15H8v-2h13a3 3 0 1 0-3-3h-2a5 5 0 1 1 5 5zm2 13a5.006 5.006 0 0 1-5-5h2a3 3 0 1 0 3-3H4v-2h19a5 5 0 0 1 0 10z"/>`),
		g.Group(children),
	)
}