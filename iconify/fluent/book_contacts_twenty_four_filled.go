package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookContactsTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.5 2A2.5 2.5 0 0 0 4 4.5v15A2.5 2.5 0 0 0 6.5 22h13.25a.75.75 0 0 0 0-1.5H6.5a1 1 0 0 1-1-1h14.25a.75.75 0 0 0 .75-.75V4.5A2.5 2.5 0 0 0 18 2H6.5Zm9 10.25v.5c0 1-1.383 1.75-3.25 1.75S9 13.75 9 12.75v-.5a.75.75 0 0 1 .75-.75h5a.75.75 0 0 1 .75.75ZM14 8.745a1.75 1.75 0 1 1-3.5 0C10.5 7.78 11.283 7 12.25 7S14 7.779 14 8.745Z"/>`),
		g.Group(children),
	)
}