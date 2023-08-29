package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HalfBrightRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M.1 9.6L3 12.5V17h4.2l2.9 2.9L13 17h4v-4.5l2.9-2.9L17 6.7V3h-3.9L10.2.1L7.2 3H3v3.7zm3.9 0c.1-3.3 2.7-5.9 6-6v12c-3.3-.1-5.9-2.7-6-6z"/>`),
		g.Group(children),
	)
}