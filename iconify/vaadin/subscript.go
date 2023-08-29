package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Subscript(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M16 15v1h-4v-1s3.3-1.6 2.6-3.2c-.5-1.1-2-.2-2-.2l-.5-.9s1.9-1.4 3.1-.2c2.4 2.3-1.4 4.5-1.4 4.5H16zM12 3H8.6L6 6L3.4 3H0l4.3 5L0 13h3.4L6 10l2.6 3H12L7.7 8z"/>`),
		g.Group(children),
	)
}