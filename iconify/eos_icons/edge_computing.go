package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EdgeComputing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.82 10.99L7 9.17V6H6V2h3v4H8v2.76l1.72 1.72A5.61 5.61 0 0 1 12 10a5.551 5.551 0 0 1 1 .09V5.91a1.5 1.5 0 1 1 1 0v4.46A5.447 5.447 0 0 1 16.29 12h2.8a1.5 1.5 0 1 1 0 1h-2.12a5.358 5.358 0 0 1 .54 1.53a3.74 3.74 0 0 1-.26 7.47H7.5a4.499 4.499 0 0 1-2.47-8.26l-1.89-1.89a1.518 1.518 0 1 1 .71-.71l2.13 2.13a4.419 4.419 0 0 1 1.03-.24a5.562 5.562 0 0 1 1.81-2.04Z"/>`),
		g.Group(children),
	)
}