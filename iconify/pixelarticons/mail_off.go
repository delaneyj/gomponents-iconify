package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h2v2H2V2zm4 4H4V4h2v2zm2 2H6V6h2v2zm2 2H8V8h2v2zm2 2h-2v-2h2v2zm2 0h-2v2h2v2h2v2h2v2h2v2h2v-2h-2v-2h-2v-2h-2v-2h-2v-2zm2-2h-2v2h2v-2zm0 0V8h2v2h-2zm-6-6h12v12h-2V6H10V4zm4 14v2H2V8h2v10h10z"/>`),
		g.Group(children),
	)
}