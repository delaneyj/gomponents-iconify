package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TreeEvergreen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m230.32 187.09l-46-59.09H208a8 8 0 0 0 6.34-12.88l-80-104a8 8 0 0 0-12.68 0l-80 104A8 8 0 0 0 48 128h23.64l-46 59.09A8 8 0 0 0 32 200h88v40a8 8 0 0 0 16 0v-40h88a8 8 0 0 0 6.32-12.91ZM48.36 184l46-59.09A8 8 0 0 0 88 112H64.25L128 29.12L191.75 112H168a8 8 0 0 0-6.31 12.91L207.64 184Z"/>`),
		g.Group(children),
	)
}