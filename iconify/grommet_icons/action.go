package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Action(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="m1 23l3-3l-3 3ZM20 4l3-3l-3 3ZM9 11l3-3l-3 3Zm4 4l3-3l-3 3ZM10 5l9 9l1-1c2-2 4.053-5 0-9s-7-2-9 0l-1 1Zm-6 6l1-1l9 9l-1 1c-2 2-5 4.087-9 0s-2-7 0-9Z"/>`),
		g.Group(children),
	)
}