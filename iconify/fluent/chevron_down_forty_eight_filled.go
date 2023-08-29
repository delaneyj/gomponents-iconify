package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronDownFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.44 15.94a1.5 1.5 0 0 1 2.12 0L24 29.378l13.44-13.44a1.5 1.5 0 0 1 2.12 2.122l-14.5 14.5a1.5 1.5 0 0 1-2.12 0l-14.5-14.5a1.5 1.5 0 0 1 0-2.122Z"/>`),
		g.Group(children),
	)
}