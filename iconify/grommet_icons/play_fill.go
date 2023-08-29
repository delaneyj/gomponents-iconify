package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="m3 22l18-10L3 2v20Zm2-3l12.6-7L5 5v14Zm2-3l7.2-4L7 8v8Zm2-3l1.8-1L9 11v2Z"/>`),
		g.Group(children),
	)
}