package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRedoTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.011 5.436a5.762 5.762 0 0 1 7.833.115L20.552 11h-6.551a1 1 0 0 0 0 2h9a1 1 0 0 0 1-1V3a1 1 0 1 0-2 0v6.618l-5.776-5.514A7.762 7.762 0 0 0 5.673 3.95c-3.324 2.991-3.439 8.166-.25 11.302l10.632 10.46a1 1 0 0 0 1.403-1.425L6.825 13.827a5.762 5.762 0 0 1 .186-8.39Z"/>`),
		g.Group(children),
	)
}