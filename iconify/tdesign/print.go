package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Print(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 2h16v5h3.5v11H19v4H5v-4H.5V7H4V2Zm2 5h12V4H6v3ZM2.5 9v7H5v-2h14v2h2.5V9h-19ZM17 16H7v4h10v-4Zm0-5.504h2.004V12.5H17v-2.004Z"/>`),
		g.Group(children),
	)
}