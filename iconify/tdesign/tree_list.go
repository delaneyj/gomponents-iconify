package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TreeList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 22a4 4 0 0 1-1-7.874V9.874A4.002 4.002 0 0 1 6 2a4 4 0 0 1 1 7.874v4.252A4.002 4.002 0 0 1 6 22Zm-2-4a2 2 0 1 0 4 0a2 2 0 0 0-4 0ZM4 6a2 2 0 1 0 4 0a2 2 0 0 0-4 0Zm8 12h7v2h-7v-2Zm0-7h10v2H12v-2Zm0-7h7v2h-7V4Z"/>`),
		g.Group(children),
	)
}