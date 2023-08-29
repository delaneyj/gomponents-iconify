package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nightlife(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 20v-2h2v-4L1 5h14l-6 9v4h2v2H5Zm.9-11h4.2l1.4-2h-7l1.4 2ZM16 20q-1.25 0-2.125-.875T13 17q0-1.25.875-2.125T16 14q.275 0 .525.038T17 14.2V5h5v3h-3v9q0 1.25-.875 2.125T16 20Z"/>`),
		g.Group(children),
	)
}