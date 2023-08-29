package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LuggageF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 4v12H7V4h6zm2 0a5 5 0 0 1 5 5v2a5 5 0 0 1-5 5V4zM5 4v12a5 5 0 0 1-5-5V9a5 5 0 0 1 5-5zm.1 0a5.002 5.002 0 0 1 9.8 0h-2.07a3.001 3.001 0 0 0-5.66 0H5.1z"/>`),
		g.Group(children),
	)
}