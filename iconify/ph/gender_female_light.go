package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GenderFemaleLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M206 96a78 78 0 1 0-84 77.75V202H88a6 6 0 0 0 0 12h34v26a6 6 0 0 0 12 0v-26h34a6 6 0 0 0 0-12h-34v-28.25A78.09 78.09 0 0 0 206 96ZM62 96a66 66 0 1 1 66 66a66.08 66.08 0 0 1-66-66Z"/>`),
		g.Group(children),
	)
}