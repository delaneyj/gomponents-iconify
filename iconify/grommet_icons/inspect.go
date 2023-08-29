package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Inspect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M5.5 21a4.5 4.5 0 1 0 0-9a4.5 4.5 0 0 0 0 9ZM1 16V6.5A4.5 4.5 0 0 1 5.5 2H6m17 14V6.5A4.5 4.5 0 0 0 18.5 2H18m.5 19a4.5 4.5 0 1 0 0-9a4.5 4.5 0 0 0 0 9ZM10 17s0-2 2-2s2 2 2 2"/>`),
		g.Group(children),
	)
}