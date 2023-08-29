package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlashThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 2a2 2 0 0 0-1.897 1.368l-4 12A2 2 0 0 0 8 18h2.415l-2.288 8.825c-.67 2.587 2.64 4.28 4.345 2.222L25.54 13.276A2 2 0 0 0 24 10h-3.114l1.987-5.298A2 2 0 0 0 21 2h-9Z"/>`),
		g.Group(children),
	)
}