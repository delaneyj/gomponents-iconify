package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Grommet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="4" d="M12 2C6.485 2 2 6.485 2 12s4.485 10 10 10s10-4.485 10-10S17.515 2 12 2Z"/>`),
		g.Group(children),
	)
}