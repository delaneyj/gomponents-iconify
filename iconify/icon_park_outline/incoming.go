package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Incoming(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linejoin="round" d="M18 13a4 4 0 1 0-8 0a4 4 0 0 0 8 0Zm19 0a4 4 0 1 0-8 0a4 4 0 0 0 8 0Z"/><path stroke-linecap="round" d="M4 13h6m8 0h11m8 0h7"/><path stroke-linecap="round" stroke-linejoin="round" d="m10 29l-6 6l6 6m-6-6h40"/></g>`),
		g.Group(children),
	)
}