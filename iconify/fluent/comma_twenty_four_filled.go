package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommaTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14.2 12.341a4 4 0 1 1 1.765-3.873c.365 1.793.507 3.948-.207 5.899c-.766 2.09-2.464 3.804-5.505 4.6a1 1 0 0 1-.506-1.934c2.46-.644 3.61-1.93 4.133-3.354c.155-.425.26-.874.32-1.338Z"/>`),
		g.Group(children),
	)
}