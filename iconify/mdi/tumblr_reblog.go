package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TumblrReblog(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M3.75 17L8 12.75V16h10v-4.5l2-2V16a2 2 0 0 1-2 2H8v3.25L3.75 17m16.5-10L16 11.25V8H6v4.5l-2 2V8a2 2 0 0 1 2-2h10V2.75L20.25 7z" fill="currentColor"/>`),
		g.Group(children),
	)
}