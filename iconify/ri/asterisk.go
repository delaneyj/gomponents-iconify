package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Asterisk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 3v7.267l6.294-3.633l1 1.732L14 11.999l6.294 3.635l-1 1.732l-6.295-3.634V21h-2v-7.268l-6.294 3.634l-1-1.732L9.998 12L3.705 8.366l1-1.732L11 10.267V3h2Z"/>`),
		g.Group(children),
	)
}