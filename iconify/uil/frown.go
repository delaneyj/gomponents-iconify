package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Frown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.36 15.33a1 1 0 0 0-.13 1.4a1 1 0 0 0 1.41.13a3.76 3.76 0 0 1 4.72 0a1 1 0 0 0 1.41-.13a1 1 0 0 0-.13-1.4a5.81 5.81 0 0 0-7.28 0ZM9 11a1 1 0 1 0-1-1a1 1 0 0 0 1 1Zm3-9a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8Zm3-11a1 1 0 1 0 1 1a1 1 0 0 0-1-1Z"/>`),
		g.Group(children),
	)
}