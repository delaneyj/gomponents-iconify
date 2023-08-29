package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SeverityCritical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m6.706.185l4.088 2.31c.437.246.706.702.706 1.195v4.62c0 .493-.269.949-.706 1.195l-4.088 2.31a1.438 1.438 0 0 1-1.412 0l-4.088-2.31A1.376 1.376 0 0 1 .5 8.31V3.69c0-.493.269-.949.706-1.195L5.294.185a1.438 1.438 0 0 1 1.412 0Z"/>`),
		g.Group(children),
	)
}