package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BastionHost(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M27.527 2.709A1.996 1.996 0 0 0 26 2h-2.69a1.5 1.5 0 0 0-1.343.83L21.382 4H20v-.5A1.5 1.5 0 0 0 18.5 2h-5A1.5 1.5 0 0 0 12 3.5V4h-1.382l-.585-1.17A1.502 1.502 0 0 0 8.69 2H6a2 2 0 0 0-1.972 2.333L4.732 8.5A3 3 0 0 0 7.69 11H8v12a3.003 3.003 0 0 0-3 3v4h22v-4a3.003 3.003 0 0 0-3-3V11h.31a3 3 0 0 0 2.958-2.5l.704-4.167a1.994 1.994 0 0 0-.445-1.624ZM25 26v2H7v-2a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1Zm-3-3H10V11h12Zm3.296-14.833A.998.998 0 0 1 24.31 9H7.69a.998.998 0 0 1-.986-.834L6 4h2.382l1 2H14V4h4v2h4.617l1.001-2H26Z"/>`),
		g.Group(children),
	)
}