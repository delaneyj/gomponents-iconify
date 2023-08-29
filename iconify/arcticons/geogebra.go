package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Geogebra(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 3a4.77 4.77 0 1 0 4.77 4.77A4.78 4.78 0 0 0 24 3ZM7.69 14.86a4.77 4.77 0 1 0 4.76 4.77a4.76 4.76 0 0 0-4.76-4.77Zm32.64 0a4.77 4.77 0 1 0 4.77 4.77a4.77 4.77 0 0 0-4.77-4.77ZM13.92 34.05a4.77 4.77 0 1 0 4.77 4.77a4.76 4.76 0 0 0-4.77-4.77Zm20.18 0a4.77 4.77 0 1 0 4.76 4.77a4.76 4.76 0 0 0-4.76-4.77Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.27 8.35a17.24 17.24 0 0 0-9.73 6.87m28.51.22a17.28 17.28 0 0 0-9.34-6.94m8.7 26.89A17.09 17.09 0 0 0 40.93 25v-.61M18 41.22a17.23 17.23 0 0 0 5.68 1a17 17 0 0 0 6.2-1.22M6.53 24.25v.72a17.16 17.16 0 0 0 3.77 10.76"/>`),
		g.Group(children),
	)
}