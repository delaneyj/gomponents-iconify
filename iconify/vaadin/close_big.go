package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloseBig(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m16 0l-1 .01L8 7L1 .01L0 0v1l7 7l-7 7v1h1l7-7l7 7h1v-1L9 8l7-7V0z"/>`),
		g.Group(children),
	)
}