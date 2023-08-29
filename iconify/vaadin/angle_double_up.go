package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngleDoubleUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M13 14v-2L8 7l-5 5v2l5-5z"/><path fill="currentColor" d="M13 9V7L8 2L3 7v2l5-5z"/>`),
		g.Group(children),
	)
}