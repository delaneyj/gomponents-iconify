package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stamper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3a3 3 0 0 0-3 3c0 3 5 7-3 7a2 2 0 0 0-2 2v2h16v-2a2 2 0 0 0-2-2c-8 0-3-4-3-7c0-2-1.34-3-3-3M6 19v2h12v-2H6Z"/>`),
		g.Group(children),
	)
}