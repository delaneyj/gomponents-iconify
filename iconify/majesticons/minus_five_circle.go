package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MinusFiveCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M23 12c0-6.075-4.925-11-11-11S1 5.925 1 12s4.925 11 11 11s11-4.925 11-11zm-9-4a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1h1.5a.5.5 0 0 1 0 1H14a1 1 0 1 0 0 2h1.5a2.5 2.5 0 0 0 0-5H15v-1h2a1 1 0 1 0 0-2h-3zm-7 3a1 1 0 1 0 0 2h3a1 1 0 1 0 0-2H7z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}