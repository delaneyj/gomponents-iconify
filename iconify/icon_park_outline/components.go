package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Components(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m17 11l7-7l7 7l-7 7l-7-7Zm13 14l7-7l7 7l-7 7l-7-7ZM17 37l7-7l7 7l-7 7l-7-7ZM4 24l7-7l7 7l-7 7l-7-7Z"/>`),
		g.Group(children),
	)
}