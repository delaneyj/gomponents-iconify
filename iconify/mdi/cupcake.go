package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cupcake(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 1.5A2.5 2.5 0 0 1 14.5 4A2.5 2.5 0 0 1 12 6.5A2.5 2.5 0 0 1 9.5 4A2.5 2.5 0 0 1 12 1.5M15.87 5C18 5 20 7 20 9c2.7 0 2.7 4 0 4H4c-2.7 0-2.7-4 0-4c0-2 2-4 4.13-4c.44 1.73 2.01 3 3.87 3c1.86 0 3.43-1.27 3.87-3M5 15h3l1 7H7l-2-7m5 0h4l-1 7h-2l-1-7m6 0h3l-2 7h-2l1-7Z"/>`),
		g.Group(children),
	)
}