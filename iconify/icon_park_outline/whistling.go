package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Whistling(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4" d="M27 11H4v8h11c-.65 1.55-1 3.21-1 5c0 7.21 5.79 13 13 13c7.2 0 13-5.79 13-13s-5.79-13-13-13Zm0 0v6m13 7h4"/>`),
		g.Group(children),
	)
}