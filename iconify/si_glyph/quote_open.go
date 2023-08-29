package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuoteOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16 14.969c.552 0 1-.435 1-.969s-.448-.969-1-.969c-2.757 0-4-1.201-4-3.907V7.906h3.441c.83 0 1.506-.605 1.506-1.438V2.512A1.51 1.51 0 0 0 15.441 1h-3.935A1.51 1.51 0 0 0 10 2.512v6.612c0 3.775 2.141 5.845 6 5.845zm-9 0c.552 0 1-.435 1-.969s-.448-.969-1-.969c-2.757 0-4-1.201-4-3.907V7.947h3.467a1.51 1.51 0 0 0 1.512-1.506V2.505A1.51 1.51 0 0 0 6.467.999H2.512A1.51 1.51 0 0 0 1 2.505v6.618c0 3.776 2.141 5.846 6 5.846z"/>`),
		g.Group(children),
	)
}