package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeechBubbleAddLtr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 1a2 2 0 0 0-2 2v16l4-4h12a2 2 0 0 0 2-2V3a2 2 0 0 0-2-2zm12 8h-4v4H9V9H5V7h4V3h2v4h4z"/>`),
		g.Group(children),
	)
}