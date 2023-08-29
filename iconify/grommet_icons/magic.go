package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Magic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="m2.5 19.5l17-17l2 2l-17 17l-2-2Zm.5.5L15 8l1 1L4 21l-1-1ZM5.5 3l-.5.5l.5.5l.5-.5l-.5-.5Zm6 0l-.5.5l.5.5l.5-.5l-.5-.5Zm-3 3l-.5.5l.5.5l.5-.5l-.5-.5Zm12 6l-.5.5l.5.5l.5-.5l-.5-.5Zm0 5l-.5.5l.5.5l.5-.5l-.5-.5Z"/>`),
		g.Group(children),
	)
}