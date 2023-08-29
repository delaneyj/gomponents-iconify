package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaveSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M8 6v9H2v2h8V8h5v18h9v-9h6v-2h-8v9h-5V6H8z"/>`),
		g.Group(children),
	)
}