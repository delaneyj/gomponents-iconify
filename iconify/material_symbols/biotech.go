package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Biotech(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21v-2h5v-2q-2.075 0-3.538-1.463T5 12q0-1.425.725-2.625t2-1.825q-.1.55.038 1.075t.437 1q-.575.4-.888 1.025T7 12q0 1.25.875 2.125T10 15h8v2h-5v2h6v2H5Zm9-8.9l-.3-.95l-.95.35l-.5-1.325q.5-.4.775-.963T13.3 8q0-1.175-.825-1.988T10.45 5.2L10 3.95l.95-.35l-.35-.9l1.9-.7l.3.95l.95-.35l2.75 7.5l-.95.35l.35.95l-1.9.7Zm-3.5-2.3q-.75 0-1.275-.525T8.7 8q0-.75.525-1.275T10.5 6.2q.75 0 1.275.525T12.3 8q0 .75-.525 1.275T10.5 9.8Z"/>`),
		g.Group(children),
	)
}