package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SetLeftRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 5c1.04 0 2.06.24 3 .68c.94-.44 1.96-.68 3-.68a7 7 0 0 1 7 7a7 7 0 0 1-7 7c-1.04 0-2.06-.24-3-.68c-.94.44-1.96.68-3 .68a7 7 0 0 1-7-7a7 7 0 0 1 7-7m0 7c0 2.22 1.21 4.16 3 5.2c1.79-1.04 3-2.98 3-5.2s-1.21-4.16-3-5.2C10.21 7.84 9 9.78 9 12Z"/>`),
		g.Group(children),
	)
}