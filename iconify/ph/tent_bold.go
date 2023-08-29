package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TentBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m255 195.13l-64-144A12 12 0 0 0 180 44H76a12 12 0 0 0-10.85 6.9a2.42 2.42 0 0 0-.12.23l-.03.17L1 195.13A12 12 0 0 0 12 212h232a12 12 0 0 0 11-16.87ZM64 112.55V188H30.46ZM88 188v-75.45L121.54 188Zm59.8 0L94.47 68h77.73l53.34 120Z"/>`),
		g.Group(children),
	)
}