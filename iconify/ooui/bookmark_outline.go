package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5 1a2 2 0 0 0-2 2v16l7-5l7 5V3a2 2 0 0 0-2-2zm10 14.25l-5-3.5l-5 3.5V3h10z"/>`),
		g.Group(children),
	)
}