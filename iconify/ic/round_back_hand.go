package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundBackHand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 16c0 4.42-3.58 8-8 8c-3.26 0-6.19-1.99-7.4-5.02l-3.03-7.61a1 1 0 0 1 1.24-1.32l.79.26c.56.18 1.02.61 1.24 1.16l1.28 3.21c.08.2.26.32.46.32H8V3.25a1.25 1.25 0 0 1 2.5 0v8.25c0 .28.22.5.5.5s.5-.22.5-.5V1.25a1.25 1.25 0 0 1 2.5 0V11.5c0 .28.22.5.5.5s.5-.22.5-.5V2.75a1.25 1.25 0 0 1 2.5 0v8.75c0 .28.22.5.5.5s.5-.22.5-.5V5.75a1.25 1.25 0 0 1 2.5 0V16z"/>`),
		g.Group(children),
	)
}