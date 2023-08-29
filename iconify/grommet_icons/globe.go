package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Globe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M2 8s3.5 1 5 2s.564 2.42 1 3c.436.58 2-1 2 2s3 1 3 4s2.5 2.5 3 1s2.233-3.134 2-5c-.233-1.866-1-3-3-3s-3.5-.5-4-2s3-2 2-5s0-4 0-4m10 11c0 6.075-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1s11 4.925 11 11Z"/>`),
		g.Group(children),
	)
}