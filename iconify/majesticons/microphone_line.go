package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrophoneLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 6a3 3 0 0 1 3-3v0a3 3 0 0 1 3 3v5a3 3 0 0 1-3 3v0a3 3 0 0 1-3-3V6zm3 12v0a7 7 0 0 1-7-7v0v-1m7 8v0a7 7 0 0 0 7-7v0v-1m-7 8v3"/>`),
		g.Group(children),
	)
}