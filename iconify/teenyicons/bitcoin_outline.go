package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BitcoinOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M3.5 1.5h5a3 3 0 1 1 0 6h-5m0-6v6m0-6H2m1.5 0V0m0 7.5h6a3 3 0 1 1 0 6h-6m0-6v6m0-6H2m1.5 6H2m1.5 0V15m4-15v1.5m0 12V15"/>`),
		g.Group(children),
	)
}