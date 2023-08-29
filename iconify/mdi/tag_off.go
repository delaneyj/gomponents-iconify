package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TagOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m18.33 8.5l4.59-4.58L21.5 2.5l-19 19l1.42 1.42l4.58-4.59l3.08 3.08A2 2 0 0 0 13 22a2 2 0 0 0 1.41-.59l7-7A2 2 0 0 0 22 13a2 2 0 0 0-.59-1.42m-15.8 3.85l9.86-9.78l-3.06-3.07A2 2 0 0 0 11 2H4a2 2 0 0 0-2 2v7a2 2 0 0 0 .59 1.41M5.5 4A1.5 1.5 0 1 1 4 5.5A1.5 1.5 0 0 1 5.5 4Z"/>`),
		g.Group(children),
	)
}