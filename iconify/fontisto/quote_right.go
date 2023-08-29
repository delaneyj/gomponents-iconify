package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuoteRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.4 24h7.2l4.8-9.6V0H0v14.4h7.2zm19.2 0h7.2l4.8-9.6V0H19.2v14.4h7.2z"/>`),
		g.Group(children),
	)
}