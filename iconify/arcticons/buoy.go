package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Buoy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.639 38.139L24 44.5l6.361-6.361m-20.5-20.5L3.5 24l6.361 6.361m20.5-20.5L24 3.5l-6.361 6.361m20.5 20.5L44.5 24l-6.361-6.361M24 16.5l7.5 7.5l-7.5 7.501l-7.5-7.5z"/>`),
		g.Group(children),
	)
}