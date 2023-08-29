package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bong(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 3v8.416c.134.059.265.123.393.193L17 8l2 2l-3.608 3.608A5 5 0 1 1 9 11.416V3h4zM8 3h6M6.1 17h9.8"/>`),
		g.Group(children),
	)
}