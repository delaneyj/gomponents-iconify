package bxs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BowlHot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 10H3a1 1 0 0 0-1 1a10 10 0 0 0 5 8.66V21a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1v-1.34A10 10 0 0 0 22 11a1 1 0 0 0-1-1zM9 9V7.93a4.51 4.51 0 0 0-1.28-3.15A2.49 2.49 0 0 1 7 3V2H5v1a4.51 4.51 0 0 0 1.28 3.17A2.49 2.49 0 0 1 7 7.93V9zm4 0V7.93a4.51 4.51 0 0 0-1.28-3.15A2.49 2.49 0 0 1 11 3V2H9v1a4.51 4.51 0 0 0 1.28 3.15A2.49 2.49 0 0 1 11 7.93V9zm4 0V7.93a4.51 4.51 0 0 0-1.28-3.15A2.49 2.49 0 0 1 15 3V2h-2v1a4.51 4.51 0 0 0 1.28 3.15A2.49 2.49 0 0 1 15 7.93V9z"/>`),
		g.Group(children),
	)
}