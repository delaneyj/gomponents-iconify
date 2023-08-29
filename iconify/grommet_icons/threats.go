package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Threats(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M9 23A7 7 0 1 0 9 9a7 7 0 0 0 0 14ZM9 6V5c0-3 2-4 4-4s4 1 4 4v3c0 1 0 3 2 3s2-2 2-3c0-2 1-2 2-2m-11 7l-6 6m0-6l6 6m-6-9V6h6v4"/>`),
		g.Group(children),
	)
}