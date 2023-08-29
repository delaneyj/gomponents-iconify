package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PowerTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15 3a1 1 0 1 0-2 0v9a1 1 0 1 0 2 0V3Zm-4.6 3.748a1 1 0 1 0-.8-1.832A11.002 11.002 0 0 0 3 15c0 6.075 4.925 11 11 11s11-4.925 11-11c0-4.512-2.717-8.388-6.6-10.084a1 1 0 0 0-.8 1.832A9.002 9.002 0 0 1 23 15a9 9 0 1 1-18 0a9.002 9.002 0 0 1 5.4-8.252Z"/>`),
		g.Group(children),
	)
}