package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronBigUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.485 15.535L12 7.05l-8.485 8.485L4.93 16.95L12 9.878l7.071 7.072l1.414-1.415Z"/>`),
		g.Group(children),
	)
}