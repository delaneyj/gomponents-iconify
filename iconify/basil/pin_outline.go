package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PinOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16.786 3.725a1.75 1.75 0 0 0-2.846.548L12.347 7.99A4.745 4.745 0 0 0 8.07 9.291l-1.71 1.71a.75.75 0 0 0 0 1.06l2.495 2.496l-5.385 5.386a.75.75 0 1 0 1.06 1.06l5.386-5.385l2.495 2.495a.75.75 0 0 0 1.061 0l1.71-1.71a4.745 4.745 0 0 0 1.302-4.277l3.716-1.592a1.75 1.75 0 0 0 .548-2.846l-3.962-3.963Zm-1.468 1.139a.25.25 0 0 1 .407-.078l3.963 3.962a.25.25 0 0 1-.079.407l-4.315 1.85a.75.75 0 0 0-.41.941a3.25 3.25 0 0 1-.763 3.396l-1.18 1.18l-4.99-4.99l1.18-1.18a3.25 3.25 0 0 1 3.396-.762a.75.75 0 0 0 .942-.41l1.85-4.316Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}