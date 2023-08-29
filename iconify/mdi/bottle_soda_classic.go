package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BottleSodaClassic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 13.77a3 3 0 0 0 0 4.46V20a2 2 0 0 1-2 2h-2a2 2 0 0 1-2-2v-1.77a3 3 0 0 0 0-4.46V9a12.28 12.28 0 0 0 1.91-5H10V2h4v2h-.91A12.28 12.28 0 0 0 15 9Z"/>`),
		g.Group(children),
	)
}