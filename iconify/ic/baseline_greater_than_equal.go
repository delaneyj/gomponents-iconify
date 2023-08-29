package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BaselineGreaterThanEqual(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m6.5 15.5l8.25-5.5L6.5 4.5l1-1.5L18 10L7.5 17z"/><path fill="currentColor" d="M18 20.998H6v-2h12z"/>`),
		g.Group(children),
	)
}