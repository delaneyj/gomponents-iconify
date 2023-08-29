package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartBoxOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 17l-.72-.66C8.7 14 7 12.46 7 10.57c0-1.54 1.21-2.75 2.75-2.75c.87 0 1.7.41 2.25 1.05c.55-.64 1.38-1.05 2.25-1.05c1.54 0 2.75 1.21 2.75 2.75c0 1.89-1.7 3.43-4.28 5.77L12 17M5 3h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2m0 2v14h14V5H5Z"/>`),
		g.Group(children),
	)
}