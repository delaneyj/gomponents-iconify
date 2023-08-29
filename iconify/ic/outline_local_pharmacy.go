package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OutlineLocalPharmacy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 5h-2.64l1.14-3.14L17.15 1l-1.46 4H3v2l2 6l-2 6v2h18v-2l-2-6l2-6V5zm-3.9 8.63L18.89 19H5.11l1.79-5.37l.21-.63l-.21-.63L5.11 7h13.78l-1.79 5.37l-.21.63l.21.63zM13 9h-2v3H8v2h3v3h2v-3h3v-2h-3z"/>`),
		g.Group(children),
	)
}