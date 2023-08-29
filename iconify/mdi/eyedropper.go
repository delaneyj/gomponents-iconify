package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eyedropper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.35 11.72l-2.13 2.13l-1.41-1.42l-7.71 7.71L3.5 22L2 20.5l1.86-4.6l7.71-7.71l-1.42-1.41l2.13-2.13l7.07 7.07M16.76 3A3 3 0 0 1 21 3a3 3 0 0 1 0 4.24l-1.92 1.92l-4.24-4.24L16.76 3M5.56 17.03L4.5 19.5l2.47-1.06L14.4 11L13 9.6l-7.44 7.43Z"/>`),
		g.Group(children),
	)
}