package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyholeCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2C6.5 2 2 6.5 2 12s4.5 10 10 10s10-4.5 10-10S17.5 2 12 2zm1.7 9c-.2.3-.4.6-.7.7V15c0 .6-.4 1-1 1s-1-.4-1-1v-3.3c-1-.6-1.3-1.8-.7-2.7S12 7.7 13 8.3c1 .5 1.3 1.7.7 2.7z"/>`),
		g.Group(children),
	)
}