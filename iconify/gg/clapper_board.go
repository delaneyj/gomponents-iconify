package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClapperBoard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m20.17 3l-.004.005A3 3 0 0 1 23 6v12a3 3 0 0 1-3 3H4a3 3 0 0 1-3-3V6a3 3 0 0 1 3-3h16.17Zm-9.694 2h6L13.09 9h-6l3.387-4ZM5.09 9l3.387-4H4a1 1 0 0 0-1 1v3h2.089ZM3 11v7a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1v-7H3Zm18-2V6a1 1 0 0 0-1-1h-1.524L15.09 9H21Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}