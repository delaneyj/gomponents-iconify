package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Blw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.857 26.893H10.448L32.333 4.5l-8.19 16.607h13.409L15.667 43.5l8.19-16.607z"/>`),
		g.Group(children),
	)
}