package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V11h-2V9h-6V3H5v18h16v2H3V1Zm12 2.414V7h3.586L15 3.414ZM8 12v8H6v-8h2Zm.5 2a2 2 0 0 1 2-2H12v2h-1.5v4H12v2h-1.5a2 2 0 0 1-2-2v-4Zm4 0a2 2 0 0 1 2-2h1a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-1a2 2 0 0 1-2-2v-4Zm5.5-2h3a2 2 0 0 1 2 2v6h-2v-6h-1v6h-2v-8Zm-3.5 2v4h1v-4h-1Z"/>`),
		g.Group(children),
	)
}