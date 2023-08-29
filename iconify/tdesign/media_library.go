package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MediaLibrary(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 2h14v2H5V2ZM3 5.5h18v2H3v-2ZM1 9h22v13H1V9Zm2 2v9h18v-9H3Zm6.75 1.469L15 15.5l-5.25 3.031V12.47Z"/>`),
		g.Group(children),
	)
}