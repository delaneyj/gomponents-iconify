package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mymail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.16 20.22L24 26.05L4.5 18.22v-5A3.68 3.68 0 0 1 8.17 9.5h31.66a3.68 3.68 0 0 1 3.67 3.67v21.66a3.68 3.68 0 0 1-3.67 3.67H8.17a3.68 3.68 0 0 1-3.67-3.67v-8.78"/>`),
		g.Group(children),
	)
}