package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowWrapTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.75 3.5a3.25 3.25 0 0 1 0 6.5h-7.5a1.75 1.75 0 1 0 0 3.5h9.19l-.22-.22a.75.75 0 1 1 1.06-1.06l1.5 1.5a.75.75 0 0 1 0 1.06l-1.5 1.5a.75.75 0 1 1-1.06-1.06l.22-.22H6.25a3.25 3.25 0 0 1 0-6.5h7.5a1.75 1.75 0 1 0 0-3.5h-10a.75.75 0 0 1 0-1.5h10Z"/>`),
		g.Group(children),
	)
}