package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HumanFemaleFemale(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.5 2a2 2 0 0 1 2 2a2 2 0 0 1-2 2a2 2 0 0 1-2-2a2 2 0 0 1 2-2M6 22v-6H3l2.6-7.6c.3-.8 1-1.4 1.9-1.4c.9 0 1.7.6 1.9 1.4L12 16l2.6-7.6c.3-.8 1-1.4 1.9-1.4c.9 0 1.7.6 1.9 1.4L21 16h-3v6h-3v-6H9v6H6M16.5 2a2 2 0 0 1 2 2a2 2 0 0 1-2 2a2 2 0 0 1-2-2a2 2 0 0 1 2-2Z"/>`),
		g.Group(children),
	)
}