package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M28.707 3.293a1 1 0 0 1 0 1.414l-24 24a1 1 0 0 1-1.414-1.414l24-24a1 1 0 0 1 1.414 0Z"/>`),
		g.Group(children),
	)
}