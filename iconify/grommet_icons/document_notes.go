package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentNotes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M4.998 9V1H19.5L23 4.5V23H4M18 1v5h5M9.75 12v6.5H9.5l-5-6.5H4v7h.25v-6.5h.25l5 6.5h.5v-7h-.25Z"/>`),
		g.Group(children),
	)
}