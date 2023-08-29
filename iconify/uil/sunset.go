package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sunset(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.66 8.34a1 1 0 0 0 .7-.29l.71-.71a1 1 0 1 0-1.41-1.41l-.66.71a1 1 0 0 0 0 1.41a1 1 0 0 0 .66.29ZM12 6a1 1 0 0 0 1-1V4a1 1 0 0 0-2 0v1a1 1 0 0 0 1 1Zm-8 6H3a1 1 0 0 0 0 2h1a1 1 0 0 0 0-2Zm1.64-3.95a1 1 0 0 0 .7.29a1 1 0 0 0 .71-.29a1 1 0 0 0 0-1.41l-.71-.71a1 1 0 0 0-1.41 1.41ZM21 12h-1a1 1 0 0 0 0 2h1a1 1 0 0 0 0-2Zm-10 7H8a1 1 0 0 0 0 2h3a1 1 0 0 0 0-2Zm7-4h-.88a5.39 5.39 0 0 0 .38-2a5.5 5.5 0 0 0-11 0a5.39 5.39 0 0 0 .38 2H6a1 1 0 0 0 0 2h12a1 1 0 0 0 0-2Zm-3.15 0h-5.7a3.44 3.44 0 0 1-.65-2a3.5 3.5 0 0 1 7 0a3.44 3.44 0 0 1-.65 2ZM16 19h-1a1 1 0 0 0 0 2h1a1 1 0 0 0 0-2Z"/>`),
		g.Group(children),
	)
}