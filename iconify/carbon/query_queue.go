package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QueryQueue(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M10 6h18v2H10zm0 6h18v2H10zm5 6h13v2H15zm-5 6h18v2H10zM4 14l7 5l-7 5V14z"/>`),
		g.Group(children),
	)
}