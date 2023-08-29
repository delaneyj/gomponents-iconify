package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OngoingConversationLtr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 0a2 2 0 0 0-2 2v18l4-4h14a2 2 0 0 0 2-2V2a2 2 0 0 0-2-2zm3 9.06a1.39 1.39 0 1 1 1.37-1.39A1.39 1.39 0 0 1 5 9.06zm5.16 0a1.39 1.39 0 1 1 1.39-1.39a1.39 1.39 0 0 1-1.42 1.39zm5.16 0a1.39 1.39 0 1 1 1.39-1.39a1.39 1.39 0 0 1-1.42 1.39z"/>`),
		g.Group(children),
	)
}