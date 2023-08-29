package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngleDoubleDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M3 2v2l5 5l5-5V2L8 7z"/><path fill="currentColor" d="M3 7v2l5 5l5-5V7l-5 5z"/>`),
		g.Group(children),
	)
}