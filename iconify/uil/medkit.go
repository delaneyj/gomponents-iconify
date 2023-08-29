package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Medkit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 17h1v1a1 1 0 0 0 2 0v-1h1a1 1 0 0 0 0-2h-1v-1a1 1 0 0 0-2 0v1h-1a1 1 0 0 0 0 2Zm9-11h-2V5a3 3 0 0 0-3-3h-4a3 3 0 0 0-3 3v1H5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3ZM9 5a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v1H9Zm11 14a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-7h16Zm0-9H4V9a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Z"/>`),
		g.Group(children),
	)
}