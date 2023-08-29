package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeechBubbleAddRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1 3v10c0 1.1.9 2 2 2h12l4 4V3c0-1.1-.9-2-2-2H3c-1.1 0-2 .9-2 2zm4 4h4V3h2v4h4v2h-4v4H9V9H5z"/>`),
		g.Group(children),
	)
}