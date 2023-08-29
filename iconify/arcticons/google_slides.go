package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleSlides(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.05 24h0v9.9H29V24Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.4 4.5a2 2 0 0 0-2 2v35a2 2 0 0 0 2 2h27.2a2 2 0 0 0 2-2v-27h-8a2 2 0 0 1-2-2v-8Zm5 14.71h17.18a1 1 0 0 1 1 1v17.16a1 1 0 0 1-1 1H15.42a1 1 0 0 1-1-1V20.21a1 1 0 0 1 1-1ZM29.61 4.5l9.99 9.99"/>`),
		g.Group(children),
	)
}