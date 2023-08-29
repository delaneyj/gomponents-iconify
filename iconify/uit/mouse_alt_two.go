package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MouseAltTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2a7 7 0 0 0-7 7v6a7.008 7.008 0 0 0 7 7a7.008 7.008 0 0 0 7-7V9a7 7 0 0 0-7-7zm6 13a6 6 0 1 1-12 0V9a6.004 6.004 0 0 1 5.5-5.974V12.5a.5.5 0 1 0 1 0V3.026A6.004 6.004 0 0 1 18 9v6z"/>`),
		g.Group(children),
	)
}