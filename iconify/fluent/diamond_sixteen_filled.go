package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiamondSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1.586 9.397a2 2 0 0 1 0-2.828L6.57 1.586a2 2 0 0 1 2.83 0l4.984 4.983a2 2 0 0 1 0 2.828L9.399 14.38a2 2 0 0 1-2.829 0L1.586 9.397Z"/>`),
		g.Group(children),
	)
}