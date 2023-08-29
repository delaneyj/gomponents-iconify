package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeechBubblesRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m0 20l4-4h10c1.1 0 2-.9 2-2v-1H5c-1.1 0-2-.9-2-2V4H2C.9 4 0 4.9 0 6zm14-10h6v6z"/><rect width="16" height="12" x="4" fill="currentColor" rx="2"/>`),
		g.Group(children),
	)
}