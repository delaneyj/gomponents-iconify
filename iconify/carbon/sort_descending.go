package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortDescending(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m18 22l1.414-1.414L23 24.172V4h2v20.172l3.586-3.586L30 22l-6 6l-6-6zM2 6h14v2H2zm4 6h10v2H6zm4 6h6v2h-6z"/>`),
		g.Group(children),
	)
}