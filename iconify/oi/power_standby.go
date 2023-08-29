package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PowerStandby(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M3 0v4h1V0H3zM1.72 1.44l-.38.31C.53 2.39 0 3.39 0 4.5C0 6.43 1.57 8 3.5 8S7 6.43 7 4.5c0-1.11-.53-2.11-1.34-2.75l-.38-.31l-.63.78l.38.31C5.61 2.99 6 3.7 6 4.5C6 5.89 4.89 7 3.5 7S1 5.89 1 4.5c0-.8.36-1.51.94-1.97l.41-.31l-.63-.78z"/>`),
		g.Group(children),
	)
}