package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlashCircleBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Zm2.018-14.636a.75.75 0 1 0-1.448-.388l-2.589 9.66a.75.75 0 0 0 1.45.388l2.587-9.66Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}