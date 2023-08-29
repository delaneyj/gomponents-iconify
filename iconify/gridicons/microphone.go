package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Microphone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 9v1a7 7 0 0 1-6 6.92V20h3v2H8v-2h3v-3.08A7 7 0 0 1 5 10V9h2v1a5 5 0 0 0 10 0V9zm-7 4a3 3 0 0 0 3-3V5a3 3 0 0 0-6 0v5a3 3 0 0 0 3 3z"/>`),
		g.Group(children),
	)
}