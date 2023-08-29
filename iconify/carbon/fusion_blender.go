package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FusionBlender(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path d="M20 12h-3V4h3V2h-8v2h3v8h-3a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h3v8h-3v2h8v-2h-3v-8h3a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2zm-8 6v-4h8v4z" fill="currentColor"/>`),
		g.Group(children),
	)
}