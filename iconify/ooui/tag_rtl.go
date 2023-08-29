package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TagRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.6 18.7c.4.4 1 .4 1.4 0L18.7 9c.2-.2.3-.4.3-.6V2c0-.6-.4-1-1-1h-6.4c-.2 0-.5.1-.6.3L1.3 11c-.4.4-.4 1-.1 1.3zM13 5c0-1.1.9-2 2-2s2 .9 2 2s-.9 2-2 2s-2-.9-2-2z"/>`),
		g.Group(children),
	)
}