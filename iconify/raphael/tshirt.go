package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tshirt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20.1 4.04a4.576 4.576 0 0 1-4.24 2.86a4.576 4.576 0 0 1-4.24-2.86L1.238 8.44l2.92 6.884l3.21-1.36V28h17.098V14.015l3.093 1.312l2.92-6.884L20.1 4.04z"/>`),
		g.Group(children),
	)
}