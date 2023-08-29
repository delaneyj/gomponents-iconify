package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Candle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 11v6h2v-6H9zm4 6h5a2 2 0 0 1 2 2v1a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2v-1a2 2 0 0 1 2-2h5V9h6v8zM2 19v1h16v-1H2zm8-10.6a3 3 0 0 1-3-3C7 4.295 8 2.495 10 0c2 2.495 3 4.295 3 5.4a3 3 0 0 1-3 3zm0-1.952c.552 0 1-.48 1-1c0-.466-.333-1.109-1-2c-.667.891-1 1.534-1 2c0 .52.448 1 1 1z"/>`),
		g.Group(children),
	)
}