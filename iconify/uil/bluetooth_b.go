package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BluetoothB(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m13.41 12l3.8-3.79a1 1 0 0 0 0-1.42l-4.5-4.5a1 1 0 0 0-.33-.21a1 1 0 0 0-.76 0a1 1 0 0 0-.54.54A1 1 0 0 0 11 3v6.59l-2.79-2.8a1 1 0 1 0-1.42 1.42l3.8 3.79l-3.8 3.79a1 1 0 1 0 1.42 1.42l2.79-2.8V21a1 1 0 0 0 .08.38a1 1 0 0 0 .54.54a.94.94 0 0 0 .76 0a1 1 0 0 0 .33-.21l4.5-4.5a1 1 0 0 0 0-1.42ZM13 5.41l2.09 2.09L13 9.59Zm0 13.18v-4.18l2.09 2.09Z"/>`),
		g.Group(children),
	)
}