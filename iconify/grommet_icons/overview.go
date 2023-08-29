package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Overview(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M18.5 21a4.5 4.5 0 1 0 0-9a4.5 4.5 0 0 0 0 9ZM10 7h4M1.5 14.5S5.5 5 6 4s1.5-1 2-1s2 0 2 2v11m-4.5 5a4.5 4.5 0 1 1 0-9a4.5 4.5 0 0 1 0 9Zm17-6.5S18.5 5 18 4s-1.5-1-2-1s-2 0-2 2v11m-4 0h4"/>`),
		g.Group(children),
	)
}