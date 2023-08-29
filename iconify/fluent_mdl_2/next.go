package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Next(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1664 256h128v1536h-128V256zM256 1792V256l1088 768l-1088 768zM384 503v1042l738-521l-738-521z"/>`),
		g.Group(children),
	)
}