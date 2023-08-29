package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RulerCombined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 4v24h24v-9H13V4H4zm2 2h5v2H8v2h3v2H8v2h3v2H8v2h3v1.586l-5 5V6zm6.414 15H14v3h2v-3h2v3h2v-3h2v3h2v-3h2v5H7.414l5-5z"/>`),
		g.Group(children),
	)
}