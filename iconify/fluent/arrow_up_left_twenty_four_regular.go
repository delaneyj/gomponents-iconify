package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpLeftTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.246 3a.75.75 0 0 1 0 1.5H5.577l15.2 15.2a.765.765 0 0 1-1.082 1.081L4.496 5.581v7.669a.75.75 0 0 1-1.5 0v-9.5a.75.75 0 0 1 .75-.75h9.5Z"/>`),
		g.Group(children),
	)
}