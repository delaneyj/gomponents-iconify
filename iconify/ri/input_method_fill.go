package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InputMethodFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 3h16a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1Zm5.869 12h4.262l.82 2h2.216L13 7h-2L6.833 17H9.05l.82-2Zm.82-2L12 9.8l1.312 3.2h-2.623Z"/>`),
		g.Group(children),
	)
}