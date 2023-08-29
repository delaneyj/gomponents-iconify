package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceEditFlipBottomFlipBottomObjectWork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 7h13a6.5 6.5 0 0 1-13 0ZM13 4.39a6.13 6.13 0 0 0-.76-1.3a6.34 6.34 0 0 0-1-1.09M1.05 4.39a6.13 6.13 0 0 1 .76-1.3A6.34 6.34 0 0 1 2.85 2M8.5.67a6.7 6.7 0 0 0-3 0"/>`),
		g.Group(children),
	)
}