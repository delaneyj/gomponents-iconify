package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WomensRoom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32c0 16.566 13.432 30 30 30s30-13.434 30-30C62 15.432 48.568 2 32 2zm-4 13c0-1 1-2 2-2h4c1 0 2 1 2 2v4c0 1-1 2-2 1.998h-4C29 21 28 20 28 19v-4zm10 21l-2-9v4l2 9h-2v11h-3V40h-2v11h-3V40h-2l2-9v-4l-2 9h-3l3-11c0-1 1-2 2-2h8c1 0 2 1 2 2l3 11h-3z"/>`),
		g.Group(children),
	)
}