package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Playcanvas(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6.115 0l-.002 3.414l5.823 3.41l-5.82 3.414l-.003 3.412l11.774-6.826zm11.77 10.35L6.113 17.174L17.887 24l-.002-3.414l-5.82-3.412l5.822-3.412z"/>`),
		g.Group(children),
	)
}