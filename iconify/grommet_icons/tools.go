package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tools(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="m11 2l11 11l-4.5 4.5l-11-11L11 2Zm5 4l1-1l2 2l-1 1m-5 5l-9 9l-2-2l9-9m-6 7l1 1"/>`),
		g.Group(children),
	)
}