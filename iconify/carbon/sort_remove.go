package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortRemove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 19.415L28.586 18L25 21.587L21.414 18L20 19.415L23.586 23L20 26.586L21.414 28L25 24.414L28.586 28L30 26.586L26.414 23L30 19.415zM10 18h6v2h-6zM2 6h14v2H2zm4 6h10v2H6z"/>`),
		g.Group(children),
	)
}