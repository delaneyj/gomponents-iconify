package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleMail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-miterlimit="10" d="M13.39 12.15v26.39H7.06A2.56 2.56 0 0 1 4.5 36V16.82m30.11-4.67v26.39h6.33A2.56 2.56 0 0 0 43.5 36V16.82"/><path fill="none" stroke="currentColor" stroke-miterlimit="10" d="M24 31.45L43.5 17v-3.6a3.94 3.94 0 0 0-6.28-3.16L24 20.06l-13.22-9.82A3.94 3.94 0 0 0 4.5 13.4V17Z"/>`),
		g.Group(children),
	)
}