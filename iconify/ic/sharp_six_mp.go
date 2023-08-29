package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpSixMp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 14h1.5v1.5H15z"/><path fill="currentColor" d="M3 3v18h18V3H3zm7 2.5h4.5V7h-3v1h3v3.5H10v-6zm2.5 13H11V14h-1v3H8.5v-3h-1v4.5H6v-6h6.5v6zM18 17h-3v1.5h-1.5v-6H18V17z"/><path fill="currentColor" d="M11.5 9H13v1.5h-1.5z"/>`),
		g.Group(children),
	)
}