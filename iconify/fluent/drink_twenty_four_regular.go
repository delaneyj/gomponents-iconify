package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DrinkTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path d="M15.72 2.22a.75.75 0 0 1 1.06 1.06L15.56 4.5h2.19c.317 0 .6.2.706.498l1.25 3.5A.75.75 0 0 1 19 9.5h-1.045l-1.587 10.05c-.21 1.465-1.46 2.45-2.97 2.45h-2.796c-1.51 0-2.76-.986-2.968-2.44L6.044 9.5H5a.75.75 0 0 1-.706-1.002l1.25-3.5A.75.75 0 0 1 6.25 4.5h7.188l2.282-2.28zm.716 7.28H7.562l1.555 9.837c.1.695.692 1.163 1.485 1.163h2.796c.793 0 1.386-.468 1.487-1.174L16.436 9.5zm.785-3.5H6.78l-.715 2h11.872l-.715-2z" fill="currentColor" fill-rule="nonzero"/>`),
		g.Group(children),
	)
}