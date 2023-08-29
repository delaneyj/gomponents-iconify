package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Superscript(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M16 5v1h-4V5s3.3-1.6 2.6-3.2c-.5-1.1-2-.2-2-.2l-.5-.9S14-.7 15.2.5C17.6 2.8 13.8 5 13.8 5H16zm-4-2H8.6L6 6L3.4 3H0l4.3 5L0 13h3.4L6 10l2.6 3H12L7.7 8z"/>`),
		g.Group(children),
	)
}