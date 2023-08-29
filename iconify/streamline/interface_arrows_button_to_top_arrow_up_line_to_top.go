package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsButtonToTopArrowUpLineToTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m.5 8.29l6.15-6.14a.48.48 0 0 1 .7 0l6.15 6.14M.5 12h13"/>`),
		g.Group(children),
	)
}