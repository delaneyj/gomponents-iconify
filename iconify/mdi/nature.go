package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nature(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 16.12a7 7 0 0 0 6.17-6.95c0-3.87-3.13-7-7-7a7 7 0 0 0-7 7c0 3.47 2.52 6.33 5.83 6.89V20H5v2h14v-2h-6v-3.88Z"/>`),
		g.Group(children),
	)
}