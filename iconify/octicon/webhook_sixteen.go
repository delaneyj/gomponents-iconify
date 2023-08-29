package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WebhookSixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M5.5 4.25a2.25 2.25 0 0 1 4.5 0a.75.75 0 0 0 1.5 0a3.75 3.75 0 1 0-6.14 2.889l-2.272 4.258a.75.75 0 0 0 1.324.706L7 7.25a.75.75 0 0 0-.309-1.015A2.25 2.25 0 0 1 5.5 4.25Z"/><path fill="currentColor" d="M7.364 3.607a.75.75 0 0 1 1.03.257l2.608 4.349a3.75 3.75 0 1 1-.628 6.785a.75.75 0 0 1 .752-1.299a2.25 2.25 0 1 0-.033-3.88a.75.75 0 0 1-1.03-.256L7.107 4.636a.75.75 0 0 1 .257-1.03Z"/><path fill="currentColor" d="M2.9 8.776A.75.75 0 0 1 2.625 9.8A2.25 2.25 0 1 0 6 11.75a.75.75 0 0 1 .75-.751h5.5a.75.75 0 0 1 0 1.5H7.425a3.751 3.751 0 1 1-5.55-3.998a.75.75 0 0 1 1.024.274Z"/>`),
		g.Group(children),
	)
}