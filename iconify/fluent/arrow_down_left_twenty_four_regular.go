package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownLeftTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.246 21.005a.75.75 0 1 0 0-1.5H5.577l15.2-15.2a.765.765 0 0 0-1.082-1.081l-15.199 15.2v-7.67a.75.75 0 0 0-1.5 0v9.5c0 .415.336.75.75.75h9.5Z"/>`),
		g.Group(children),
	)
}