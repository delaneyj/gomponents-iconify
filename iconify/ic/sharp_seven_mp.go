package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpSevenMp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 14h1.5v1.5H15z"/><path fill="currentColor" d="M3 3v18h18V3H3zm9.5 15.5H11V14h-1v3H8.5v-3h-1v4.5H6v-6h6.5v6zm-1.25-7L12.62 7H10V5.5h4.87l-1.87 6h-1.75zM18 17h-3v1.5h-1.5v-6H18V17z"/>`),
		g.Group(children),
	)
}