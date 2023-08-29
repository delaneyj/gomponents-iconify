package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IndentLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 20h2V4h-2v16zM2 11h10.172l-2.086-2.086L11.5 7.5L16 12l-4.5 4.5l-1.414-1.414L12.172 13H2v-2z"/>`),
		g.Group(children),
	)
}