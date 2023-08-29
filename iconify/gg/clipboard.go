package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clipboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M8 11a1 1 0 1 0 0 2h7.96a1 1 0 1 0 0-2H8Zm.04 4.066a1 1 0 1 0 0 2H16a1 1 0 1 0 0-2H8.04Z"/><path fill-rule="evenodd" d="M5 3a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2H5Zm2 2H5v14h14V5h-2v1a3 3 0 0 1-3 3h-4a3 3 0 0 1-3-3V5Zm2 0v1a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1V5H9Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}