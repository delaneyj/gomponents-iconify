package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 5a1 1 0 0 1 1-1h18a1 1 0 0 1 1 1v4a1 1 0 1 1-2 0V6h-7v20h2a1 1 0 1 1 0 2h-6a1 1 0 1 1 0-2h2V6H8v3a1 1 0 0 1-2 0V5Z"/>`),
		g.Group(children),
	)
}