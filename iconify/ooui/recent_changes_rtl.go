package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RecentChangesRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M19 3H3v2h16zm0 6h-7v2h7zm0 6H8v2h11zM8.8 11.9l1-1.1a.6.6 0 0 0 0-.8L8 8.2a.6.6 0 0 0-.8 0l-1 1L8.7 12zm-3.3-2L0 15.3V18h2.6l5.6-5.5l-2.7-2.7Z"/>`),
		g.Group(children),
	)
}