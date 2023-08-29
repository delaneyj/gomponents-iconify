package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwoThirty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2m2 33.445V58h-4V35.445a3.963 3.963 0 0 1-.944-.758l-3.029.813L25 31.666l3.201-.859A3.983 3.983 0 0 1 30 28.555V26h4v2.555c.251.146.484.313.697.509L47.975 25.5L49 29.336L35.905 32.85A3.986 3.986 0 0 1 34 35.445"/><circle cx="32" cy="32" r="3" fill="currentColor"/>`),
		g.Group(children),
	)
}