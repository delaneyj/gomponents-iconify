package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DaFit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.523 38.141A14.836 14.836 0 0 1 24 42.5c-8.22 0-14.882-6.663-14.882-14.882S15.78 12.736 24 12.736s14.882 6.663 14.882 14.882m0 0V42.5m0-37v13.11"/>`),
		g.Group(children),
	)
}