package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Emby(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 2L6 7l1 1l-5 5l5 5l1-1l5 5l5-5l-1-1l5-5l-5-5l-1 1l-5-5m-1 6.5l6 3.5l-6 3.5v-7Z"/>`),
		g.Group(children),
	)
}