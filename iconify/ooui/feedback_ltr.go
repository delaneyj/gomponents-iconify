package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FeedbackLtr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M19 16L2 12a3.83 3.83 0 0 1-1-2.5A3.83 3.83 0 0 1 2 7l17-4z"/><rect width="4" height="8" x="4" y="9" fill="currentColor" rx="2"/>`),
		g.Group(children),
	)
}