package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dashboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M4 0C1.8 0 0 1.8 0 4s1.8 4 4 4s4-1.8 4-4s-1.8-4-4-4zm0 1c1.66 0 3 1.34 3 3S5.66 7 4 7S1 5.66 1 4s1.34-3 3-3zm0 1c-.28 0-.5.22-.5.5s.22.5.5.5s.5-.22.5-.5S4.28 2 4 2zM2.34 3a.5.5 0 0 0-.19.84l.91.91C3.04 4.83 3 4.91 3 5c0 .55.45 1 1 1s1-.45 1-1s-.45-1-1-1c-.09 0-.17.04-.25.06l-.91-.91a.5.5 0 0 0-.44-.16a.5.5 0 0 0-.06 0zM5.5 3c-.28 0-.5.22-.5.5s.22.5.5.5s.5-.22.5-.5s-.22-.5-.5-.5z"/>`),
		g.Group(children),
	)
}