package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 2h14v12h-2V4H8V2zM2 8h2v6H2V8zm2 6h2v2H4v-2zm4 2H6v2h2v2h2v2h4v-2h-4v-2H8v-2zm10 0h-2v2h2v2h2v2h2v-2h-2v-2h-2v-2zM4 2H2v2h2v2h2v2h2v2h2v2h2v2h2v2h2v-2h-2v-2h-2v-2h-2V8H8V6H6V4H4V2z"/>`),
		g.Group(children),
	)
}