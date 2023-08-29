package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuickCalculation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.189 19.88a1.428 1.428 0 0 0-1.097-.512H8.215a1.643 1.643 0 0 0-1.62 1.36l-1.097 6.224a1.431 1.431 0 0 0 1.41 1.68h32.877a1.643 1.643 0 0 0 1.62-1.36l1.097-6.224a1.429 1.429 0 0 0-.313-1.169Zm-16.715-3.7a6.543 6.543 0 0 0 6.178-5.247a4.638 4.638 0 0 0-.947-3.86A4.516 4.516 0 0 0 27.167 5.5a6.543 6.543 0 0 0-6.178 5.246a4.638 4.638 0 0 0 .947 3.86a4.516 4.516 0 0 0 3.538 1.573Zm-2.948 15.64a6.543 6.543 0 0 0-6.178 5.247a4.638 4.638 0 0 0 .947 3.86a4.515 4.515 0 0 0 3.538 1.573a6.543 6.543 0 0 0 6.178-5.246a4.638 4.638 0 0 0-.947-3.86a4.516 4.516 0 0 0-3.538-1.573Z"/>`),
		g.Group(children),
	)
}