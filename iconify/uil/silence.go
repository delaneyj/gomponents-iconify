package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Silence(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 11a1 1 0 1 0-1-1a1 1 0 0 0 1 1Zm6 0a1 1 0 1 0-1-1a1 1 0 0 0 1 1Zm1 3a1 1 0 0 0-2 0h-1a1 1 0 0 0-2 0h-1a1 1 0 0 0-2 0a1 1 0 0 0 0 2a1 1 0 0 0 2 0h1a1 1 0 0 0 2 0h1a1 1 0 0 0 2 0a1 1 0 0 0 0-2ZM12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8Z"/>`),
		g.Group(children),
	)
}