package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ResizeVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M904.102 304.102L600 0L295.898 304.102h203.539v591.797H295.898L600 1200l304.102-304.102h-203.54V304.102h203.54z"/>`),
		g.Group(children),
	)
}