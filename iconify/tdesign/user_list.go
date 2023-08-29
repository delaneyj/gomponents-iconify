package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.5 4a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7ZM6 7.5a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0ZM8 16a4 4 0 0 0-4 4h7.8v2H2v-2a6 6 0 0 1 6-6h3.75v2H8Zm5.75-2h8.5v2h-8.5v-2Zm0 3h8.5v2h-8.5v-2Zm0 3h8.5v2h-8.5v-2Z"/>`),
		g.Group(children),
	)
}