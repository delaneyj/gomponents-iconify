package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoInputComponent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 2a1 1 0 0 0-1-1a1 1 0 0 0-1 1v4H1v6h6V6H5V2m4 14c0 1.3.84 2.4 2 2.82V23h2v-4.18c1.16-.41 2-1.51 2-2.82v-2H9v2m-8 0c0 1.3.84 2.4 2 2.82V23h2v-4.18C6.16 18.4 7 17.3 7 16v-2H1v2M21 6V2a1 1 0 0 0-1-1a1 1 0 0 0-1 1v4h-2v6h6V6h-2m-8-4a1 1 0 0 0-1-1a1 1 0 0 0-1 1v4H9v6h6V6h-2V2m4 14c0 1.3.84 2.4 2 2.82V23h2v-4.18c1.16-.41 2-1.51 2-2.82v-2h-6v2Z"/>`),
		g.Group(children),
	)
}