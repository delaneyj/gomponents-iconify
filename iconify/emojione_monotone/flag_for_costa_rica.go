package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForCostaRica(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2m-17.755 8.366h35.511a28.155 28.155 0 0 1 7.628 9.834H6.616a28.167 28.167 0 0 1 7.629-9.834m35.511 43.267H14.244A28.158 28.158 0 0 1 6.617 43.8h50.768a28.184 28.184 0 0 1-7.629 9.833"/>`),
		g.Group(children),
	)
}