package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Myxl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m31.551 34.467l-14.598-16.99H8.122l7.298 8.495l-7.298 8.495h8.831l2.883-3.355l2.883 3.355h8.832z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.645 31.085v-3.094h7.021l4.082-6.645H28.645V10.243L22 14.324v9.027"/>`),
		g.Group(children),
	)
}