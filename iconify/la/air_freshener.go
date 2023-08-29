package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirFreshener(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 2c-1.654 0-3 1.346-3 3c0 .854.363 1.621.938 2.168L9.056 14h2.076l-4 6H15v2H9v6h14v-6h-6v-2h7.867l-4-6h2.076l-4.879-6.832A2.986 2.986 0 0 0 19 5c0-1.654-1.346-3-3-3zm0 2c.552 0 1 .449 1 1a.999.999 0 0 1-.59.906l-.076.028A.979.979 0 0 1 16 6a.979.979 0 0 1-.334-.066l-.076-.028A.999.999 0 0 1 15 5c0-.551.448-1 1-1zm-.193 3.992c.038.003.076 0 .115 0c.026 0 .051.008.078.008c.027 0 .052-.007.078-.008c.04 0 .078.003.117 0L19.057 12h-1.924l4 6H10.867l4-6h-1.924l2.864-4.008zM11 24h10v2H11v-2z"/>`),
		g.Group(children),
	)
}