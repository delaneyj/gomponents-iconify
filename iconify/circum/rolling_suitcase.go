package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RollingSuitcase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.25 7.49H15V3.56a1.5 1.5 0 0 0-1.5-1.5h-3A1.511 1.511 0 0 0 9 3.56v3.93H7.75a2.5 2.5 0 0 0-2.5 2.5v8.44a2.5 2.5 0 0 0 2.5 2.5h.5v.01a1 1 0 0 0 2 0v-.01h3.5v.01a1 1 0 0 0 2 0v-.01h.5a2.5 2.5 0 0 0 2.5-2.5V9.99a2.5 2.5 0 0 0-2.5-2.5ZM10 3.56a.508.508 0 0 1 .5-.5h3a.5.5 0 0 1 .5.5V7.5h-4Zm7.75 14.87a1.5 1.5 0 0 1-1.5 1.5h-8.5a1.5 1.5 0 0 1-1.5-1.5V9.99a1.511 1.511 0 0 1 1.5-1.5h8.5a1.5 1.5 0 0 1 1.5 1.5Z"/>`),
		g.Group(children),
	)
}