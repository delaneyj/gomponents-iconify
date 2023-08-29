package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeechBubblesLtr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17 4v7a2 2 0 0 1-2 2H4v1a2 2 0 0 0 2 2h10l4 4V6a2 2 0 0 0-2-2zM6 10H0v6z"/><rect width="16" height="12" fill="currentColor" rx="2"/>`),
		g.Group(children),
	)
}