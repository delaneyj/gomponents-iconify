package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WbAuto(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.825 15h1.2l.65-1.8h2.8l.65 1.8h1.2l-2.6-7h-1.3l-2.6 7Zm2.2-2.8l1-2.9h.1l1 2.9h-2.1Zm1.05 6.8q-2.925 0-4.962-2.038T1.075 12q0-2.925 2.038-4.963T8.075 5q1.875 0 3.463.925t2.537 2.525l-.1-.45h1.2l1.2 5.1h.05l1.45-5.1h1.1l1.45 5.1h.1l1.2-5.1h1.2l-1.85 7h-1.15l-1.5-5.25h-.05L16.925 15h-1.15l-.7-2.9q0 2.875-2.05 4.888T8.075 19Z"/>`),
		g.Group(children),
	)
}