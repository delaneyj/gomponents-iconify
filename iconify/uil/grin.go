package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Grin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 11a1 1 0 1 0-1-1a1 1 0 0 0 1 1Zm3-9a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8Zm3-7H9a1 1 0 0 0-1 1a4 4 0 0 0 8 0a1 1 0 0 0-1-1Zm-3 3a2 2 0 0 1-1.73-1h3.46A2 2 0 0 1 12 16Zm3-7a1 1 0 1 0 1 1a1 1 0 0 0-1-1Z"/>`),
		g.Group(children),
	)
}