package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpBeenhere(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.01 1L3 17l9 6l8.99-6L21 1H3.01zM10 16l-5-5l1.41-1.42L10 13.17l7.59-7.59L19 7l-9 9z"/>`),
		g.Group(children),
	)
}