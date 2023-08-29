package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsButtonToLeftArrowLineToLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.29.5L2.15 6.65a.48.48 0 0 0 0 .7l6.14 6.15M12 .5v13"/>`),
		g.Group(children),
	)
}