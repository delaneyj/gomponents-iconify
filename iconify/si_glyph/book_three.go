package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m4.048 13.774l2.014 2.122l8.847-2.065V1.965l-.977-.581l-7.87 2.577l-1.065-1.04l7.7-2.285l-1.148-.625l-7.487 2.122l-.014 11.641z"/>`),
		g.Group(children),
	)
}