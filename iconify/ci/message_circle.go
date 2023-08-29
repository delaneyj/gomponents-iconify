package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m8.051 17.828l.654.35a6.96 6.96 0 0 0 3.292.822H12a7 7 0 1 0-7-7v.003a6.96 6.96 0 0 0 .822 3.292l.35.654l-.538 2.417l2.417-.538ZM3 21l1.058-4.762A9 9 0 0 1 12 3a9 9 0 0 1 9 9a9 9 0 0 1-13.238 7.942L3 21Z"/>`),
		g.Group(children),
	)
}