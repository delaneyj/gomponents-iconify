package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Success(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 20a10 10 0 0 1 0-20a10 10 0 1 1 0 20Zm-2-5l9-8.5L15.5 5L8 12L4.5 8.5L3 10l5 5Z"/>`),
		g.Group(children),
	)
}