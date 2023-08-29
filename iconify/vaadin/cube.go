package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cube(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 0L0 2v10l8 4l8-4V2L8 0zm6.4 2.6L8.5 4.8L1.9 2.6L8 1l6.4 1.6zM1 11.4V3.3l7 2.4v9.2l-7-3.5z"/>`),
		g.Group(children),
	)
}