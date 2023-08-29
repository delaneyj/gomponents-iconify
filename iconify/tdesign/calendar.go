package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Calendar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 4V1.5h2V4h8V1.5h2V4h4v18H2V4h4ZM4 6v3h16V6H4Zm16 5H4v9h16v-9Z"/>`),
		g.Group(children),
	)
}