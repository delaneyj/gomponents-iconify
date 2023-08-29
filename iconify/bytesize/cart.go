package bytesize

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 6h24l-3 13H9m18 4H10L5 2H2"/><circle cx="25" cy="27" r="2"/><circle cx="12" cy="27" r="2"/></g>`),
		g.Group(children),
	)
}