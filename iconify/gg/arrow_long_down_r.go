package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLongDownR(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16.242 4.641L11.999.4L7.756 4.64L11 7.886l.013 11.9l-1.845-1.834l-1.41 1.418l4.255 4.231l4.23-4.254l-1.417-1.41l-1.813 1.822L13 7.883l3.242-3.242Zm-5.657 0l1.414-1.414l1.414 1.414L12 6.056L10.585 4.64Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}