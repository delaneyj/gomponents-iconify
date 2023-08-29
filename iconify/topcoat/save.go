package topcoat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Save(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 42 42"),
		g.Raw(`<path fill="currentColor" d="M24.531 1.5H17.5c-2.53 0-3 .566-3 3.016V15.5H5l16.531 16.969L38 15.5H27.5v-11c0-2.453-.453-3-2.969-3zM11 36.5l-2.63-6H3l3.94 9.62c.35 1.05 1.091 1.38 3.123 1.38H33c1.969 0 2.38-.3 2.729-1.35l3.94-9.621h-5.13L32 36.5H11z"/>`),
		g.Group(children),
	)
}