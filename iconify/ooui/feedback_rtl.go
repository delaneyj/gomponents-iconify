package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FeedbackRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18 7c.6.7 1 1.6 1 2.5c0 .9-.4 1.8-1 2.5L1 16V3z"/><rect width="4" height="8" x="12" y="9" fill="currentColor" rx="2"/>`),
		g.Group(children),
	)
}