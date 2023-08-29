package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextWidth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M8 6v2h7v14h2V8h7V6zm2 15.5L5.625 25L10 28.5V26h12v2.5l4.375-3.5L22 21.5V24H10z"/>`),
		g.Group(children),
	)
}