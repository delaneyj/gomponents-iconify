package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Celsius(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 19h-6a3 3 0 0 1-3-3V8a3 3 0 0 1 3-3h6a1 1 0 0 0 0-2h-6a5 5 0 0 0-5 5v8a5 5 0 0 0 5 5h6a1 1 0 0 0 0-2ZM5 3a3 3 0 1 0 3 3a3 3 0 0 0-3-3Zm0 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1Z"/>`),
		g.Group(children),
	)
}