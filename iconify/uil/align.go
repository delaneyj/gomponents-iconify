package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Align(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 15h10a1 1 0 0 0 0-2H3a1 1 0 0 0 0 2Zm0-4h10a1 1 0 0 0 0-2H3a1 1 0 0 0 0 2Zm0-4h18a1 1 0 0 0 0-2H3a1 1 0 0 0 0 2Zm18 10H3a1 1 0 0 0 0 2h18a1 1 0 0 0 0-2Zm-1.36-7.43a1 1 0 1 0-1.28 1.53l1.08.9l-1.08.9a1 1 0 0 0-.13 1.41a1 1 0 0 0 .77.36a1 1 0 0 0 .64-.24l2-1.66a1 1 0 0 0 0-1.54Z"/>`),
		g.Group(children),
	)
}