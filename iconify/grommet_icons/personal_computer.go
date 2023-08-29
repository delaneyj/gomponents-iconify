package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonalComputer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M3 18h18V5a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v13Zm-1 2h20c1 0 1-1 1-1H1s0 1 1 1Z"/>`),
		g.Group(children),
	)
}