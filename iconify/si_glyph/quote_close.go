package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuoteClose(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 14.969c-.553 0-1-.435-1-.969s.447-.969 1-.969c2.757 0 4-1.201 4-3.907V7.978H2.559a1.51 1.51 0 0 1-1.506-1.511V2.511A1.51 1.51 0 0 1 2.559.999h3.935C7.324.999 8 1.677 8 2.511v6.612c0 3.775-2.141 5.846-6 5.846zm9 0c-.553 0-1-.435-1-.969s.447-.969 1-.969c2.757 0 4-1.201 4-3.907V7.947h-3.467a1.51 1.51 0 0 1-1.512-1.506V2.505c0-.83.678-1.506 1.512-1.506h3.955A1.51 1.51 0 0 1 17 2.505v6.618c0 3.775-2.141 5.846-6 5.846z"/>`),
		g.Group(children),
	)
}