package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11 17.9A5.002 5.002 0 0 1 7 13V7a5 5 0 0 1 10 0v6a5.002 5.002 0 0 1-4 4.9V21a1 1 0 1 1-2 0v-3.1ZM12 4a3 3 0 0 1 3 3v6a3.001 3.001 0 0 1-2 2.83V11a1 1 0 1 0-2 0v4.83A3.001 3.001 0 0 1 9 13V7a3 3 0 0 1 3-3Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}