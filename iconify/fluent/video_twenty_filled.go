package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13 6.5A2.5 2.5 0 0 0 10.5 4h-6A2.5 2.5 0 0 0 2 6.5v7A2.5 2.5 0 0 0 4.5 16h6a2.5 2.5 0 0 0 2.5-2.5v-7Zm1 1.43v4.152l2.764 2.35A.75.75 0 0 0 18 13.86V6.193a.75.75 0 0 0-1.23-.575L14 7.93Z"/>`),
		g.Group(children),
	)
}