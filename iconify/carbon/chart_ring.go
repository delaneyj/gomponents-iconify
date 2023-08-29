package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartRing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 2a14 14 0 1 0 14 14A14.016 14.016 0 0 0 16 2Zm11.95 13.001h-6.04a5.982 5.982 0 0 0-9.38-3.885l-4.27-4.27A11.978 11.978 0 0 1 27.95 15ZM16 20a4 4 0 1 1 4-4a4.005 4.005 0 0 1-4 4ZM4 16a11.937 11.937 0 0 1 2.846-7.74l4.27 4.27A5.984 5.984 0 0 0 15 21.91v6.04A12.01 12.01 0 0 1 4 16Zm13 11.95v-6.04A6.007 6.007 0 0 0 21.91 17h6.04A12.008 12.008 0 0 1 17 27.95Z"/>`),
		g.Group(children),
	)
}