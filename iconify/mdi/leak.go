package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Leak(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 3H3v3a3 3 0 0 0 3-3m8 0h-2a9 9 0 0 1-9 9v2c6.08 0 11-4.93 11-11m-4 0H8a5 5 0 0 1-5 5v2a7 7 0 0 0 7-7m0 18h2a9 9 0 0 1 9-9v-2a11 11 0 0 0-11 11m8 0h3v-3a3 3 0 0 0-3 3m-4 0h2a5 5 0 0 1 5-5v-2a7 7 0 0 0-7 7Z"/>`),
		g.Group(children),
	)
}