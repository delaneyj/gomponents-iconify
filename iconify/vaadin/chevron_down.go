package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m8 13.1l-8-8l2.1-2.2L8 8.8l5.9-5.9L16 5.1z"/>`),
		g.Group(children),
	)
}