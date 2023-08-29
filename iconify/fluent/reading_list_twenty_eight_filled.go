package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReadingListTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 7.5a1.5 1.5 0 0 1 2.75-.83a1 1 0 0 0 1.663-1.11A3.5 3.5 0 1 0 5.483 11H21a1 1 0 1 0 0-2H5.5A1.5 1.5 0 0 1 4 7.5ZM12 5a1 1 0 1 0 0 2h13a1 1 0 1 0 0-2H12Zm-5 8a1 1 0 1 0 0 2h18a1 1 0 1 0 0-2H7Zm-5 5a1 1 0 0 1 1-1h18a1 1 0 1 1 0 2H3a1 1 0 0 1-1-1Zm5 3a1 1 0 1 0 0 2h18a1 1 0 1 0 0-2H7Z"/>`),
		g.Group(children),
	)
}