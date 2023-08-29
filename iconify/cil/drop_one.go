package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DropOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M18.498 10.344L12.631 1.18a.751.751 0 0 0-1.262-.003l-.002.003L5.5 10.344a7.935 7.935 0 0 0-1.559 4.743v.022v-.001c0 4.443 3.615 8.058 8.058 8.058s8.058-3.615 8.058-8.058v-.021a7.957 7.957 0 0 0-1.574-4.764l.015.021zM12 21.666a6.566 6.566 0 0 1-6.558-6.558v-.018c0-1.46.481-2.807 1.293-3.893l-.012.017l.029-.041l5.249-8.198l5.248 8.198l.029.041a6.456 6.456 0 0 1 1.281 3.876v.019v-.001a6.566 6.566 0 0 1-6.558 6.558z" fill="currentColor"/>`),
		g.Group(children),
	)
}