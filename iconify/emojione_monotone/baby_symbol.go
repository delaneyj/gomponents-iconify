package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BabySymbol(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32c0 16.566 13.432 30 30 30s30-13.434 30-30C62 15.432 48.568 2 32 2m-5 12.5c0-1.25 1.25-2.5 2.5-2.5h5c1.25 0 2.5 1.25 2.5 2.5v5c0 1.25-1.25 2.502-2.5 2.5h-5c-1.25.002-2.5-1.25-2.5-2.5v-5M29 51l-3 1l-3-7l3-7l5 3l-4 4l2 6m9 1l-3-1l2-6l-4-4l5-3l3 7l-3 7m0-24v6c0 2-1 3-3 3h-6c-2 0-3-1-3-3v-6l-7-5l1-2l7.946 3H36l8-3l1 2l-7 5"/>`),
		g.Group(children),
	)
}