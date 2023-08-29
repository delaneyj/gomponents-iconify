package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatColorResetRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.25 16.4L8.4 5.55l2.55-2.525q.225-.225.5-.325t.55-.1q.275 0 .55.113t.5.312l4.6 4.525q1.1 1.05 1.725 2.488T20 13.1q0 .9-.2 1.725t-.55 1.575ZM12 21q-3.325 0-5.663-2.313T4 13.1q0-1.275.4-2.45T5.6 8.4L2.1 4.9q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l17 17q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275l-2.4-2.4q-1.025.725-2.2 1.113T12 21Z"/>`),
		g.Group(children),
	)
}