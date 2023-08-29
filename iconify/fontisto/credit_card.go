package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 21.5V12h36v9.5a2.507 2.507 0 0 1-2.5 2.5H2.466a2.4 2.4 0 0 1-1.731-.734l-.001-.001a2.401 2.401 0 0 1-.735-1.731v-.036v.002zM10 18v2h6v-2zm-6 0v2h4v-2zM33.5 0A2.507 2.507 0 0 1 36 2.5V6H0V2.466A2.4 2.4 0 0 1 .734.735L.735.734a2.401 2.401 0 0 1 1.731-.735h.036H2.5z"/>`),
		g.Group(children),
	)
}