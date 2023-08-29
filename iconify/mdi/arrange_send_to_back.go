package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrangeSendToBack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h9v9H2V2m7 2H4v5h5V4m13 9v9h-9v-9h9m-7 7h5v-5h-5v5m1-12v3h-3V8h3m-5 8H8v-3h3v3Z"/>`),
		g.Group(children),
	)
}