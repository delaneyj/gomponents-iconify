package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Drive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M3 14h10a3 3 0 0 0 3-3H0a3 3 0 0 0 3 3zm10-2h1v1h-1v-1zm2-10H1l-1 8h16z"/>`),
		g.Group(children),
	)
}