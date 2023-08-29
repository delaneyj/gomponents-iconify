package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LanguageLua(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.5 5A8.5 8.5 0 0 0 2 13.5a8.5 8.5 0 0 0 8.5 8.5a8.5 8.5 0 0 0 8.5-8.5A8.5 8.5 0 0 0 10.5 5m3 8a2.5 2.5 0 0 1-2.5-2.5A2.5 2.5 0 0 1 13.5 8a2.5 2.5 0 0 1 2.5 2.5a2.5 2.5 0 0 1-2.5 2.5m6-11A2.5 2.5 0 0 0 17 4.5A2.5 2.5 0 0 0 19.5 7A2.5 2.5 0 0 0 22 4.5A2.5 2.5 0 0 0 19.5 2"/>`),
		g.Group(children),
	)
}