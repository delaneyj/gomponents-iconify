package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataDisplay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 18H3v2h19v-2Zm-2-2a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2h-4v2h4v3h-3.5v2H20v3h-4v2h4Zm-5-2h-4v-3h2a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2H9v2h4v3h-2a2 2 0 0 0-2 2v5h6v-2Zm-7 0H6.5V6a2 2 0 0 0-2-2H3v2h1.5v8H3v2h5v-2Z"/>`),
		g.Group(children),
	)
}