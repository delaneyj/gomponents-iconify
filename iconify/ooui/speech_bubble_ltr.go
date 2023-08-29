package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeechBubbleLtr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 14H0v6z"/><rect width="20" height="16" fill="currentColor" rx="2"/>`),
		g.Group(children),
	)
}