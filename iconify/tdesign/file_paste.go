package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilePaste(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V23H3V1Zm16 8h-6V3H5v18h14V9Zm-4-5.583V7h3.585L15 3.417ZM7 10h2.6v2H9v.6H7V10Zm3.4 0H13v2.6h-2V12h-.6v-2ZM9 14v-.6H7V16h2.6v-2H9Zm2 0h6v6h-6v-6Zm2 2v2h2v-2h-2Z"/>`),
		g.Group(children),
	)
}