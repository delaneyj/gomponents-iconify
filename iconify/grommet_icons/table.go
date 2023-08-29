package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Table(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M8 5v18m8-18v18M1 11h22M1 5h22M1 17h22M1 1h22v22H1V1Z"/>`),
		g.Group(children),
	)
}