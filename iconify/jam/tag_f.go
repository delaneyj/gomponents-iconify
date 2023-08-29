package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TagF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17.014.73l3.448 3.447a2 2 0 0 1 .463 2.103l-1.763 4.812a1 1 0 0 1-.232.363l-8.323 8.323a2 2 0 0 1-2.829 0l-6.364-6.364a2 2 0 0 1 0-2.828l8.333-8.333a1 1 0 0 1 .364-.232L14.913.266a2 2 0 0 1 2.101.464zM13.79 7.404a1.5 1.5 0 1 0 2.12-2.122a1.5 1.5 0 0 0-2.12 2.122z"/>`),
		g.Group(children),
	)
}