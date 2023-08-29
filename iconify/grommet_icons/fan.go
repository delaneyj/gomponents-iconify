package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M12 15a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm3-3c4 3 5 7 5 7m-8 4c6.075 0 11-4.925 11-11S18.075 1 12 1S1 5.925 1 12s4.925 11 11 11Zm0-14c3-5 7-6 7-6m-7 12c-3 4-7 5-7 5m4-8C5 9 4 5 4 5"/>`),
		g.Group(children),
	)
}