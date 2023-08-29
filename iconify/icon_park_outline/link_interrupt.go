package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkInterrupt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M27 14h15a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2h-4M11 14H6a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h15M14 6l20 34m-2-17h4m-24 0h4"/>`),
		g.Group(children),
	)
}