package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CeilingLightMultiple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 11h3V6h2v5h3l4 8H6l4-8m6 9c0 1.11-.89 2-2 2s-2-.89-2-2h4m-7.79-9.89L8.76 9H11V2H9v5H6l-4 8h3.76l2.45-4.89Z"/>`),
		g.Group(children),
	)
}