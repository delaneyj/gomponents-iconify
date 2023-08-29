package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarHalf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2L9.19 8.62L2 9.24l5.45 4.73L5.82 21L12 17.27V2Z"/>`),
		g.Group(children),
	)
}