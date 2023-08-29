package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M5 0h11v3H5V0zM1 4h15v3H1V4zm2 4h13v3H3V8zm-3 4h16v3H0v-3z"/>`),
		g.Group(children),
	)
}