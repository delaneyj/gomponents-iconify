package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m8 2.9l8 8l-2.1 2.2L8 7.2l-5.9 5.9L0 10.9z"/>`),
		g.Group(children),
	)
}