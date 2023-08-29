package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageFlash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 2H2v20h2V4h16v10h2V2h-2zM10 16H6v2H4v2h2v-2h4v-2zm6-4h2v4h4v2h-2v2h-2v2h-2v-4h-4v-2h2v-2h2v-2z"/>`),
		g.Group(children),
	)
}