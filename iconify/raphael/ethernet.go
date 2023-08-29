package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ethernet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M22.5 8.5v3.168l3.832 3.832l-3.832 3.832V22.5l7-7l-7-7zm-14 14v-3.168L4.667 15.5L8.5 11.668V8.5l-7 7l7 7zm7-8.4a1.68 1.68 0 1 0 0 3.356a1.68 1.68 0 0 0 0-3.354zm-5.04 0a1.68 1.68 0 1 0 0 3.362a1.681 1.681 0 0 0 0-3.36zm10.08 0a1.68 1.68 0 1 0 0 3.357a1.68 1.68 0 0 0 0-3.355z"/>`),
		g.Group(children),
	)
}