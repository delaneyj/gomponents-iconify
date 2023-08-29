package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoNotStepSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m18.5 15.675l-1.4-1.4L21.575 9.8L23 11.2l-4.5 4.475Zm-2.1-2.125L9.3 6.475L13.8 2l7.075 7.075L16.4 13.55Zm3.375 9.05l-5.95-5.925L10.5 20H1v-3.95q.85-.35 1.675-.713t1.625-.862L6.025 16.2l.7-.7l-1.65-1.65l.388-.388q.187-.187.387-.412L7.5 14.7l.7-.7l-1.775-1.775q.15-.275.263-.55t.187-.575l2.1 2.1l.7-.7L1.4 4.225L2.8 2.8l18.375 18.4l-1.4 1.4Z"/>`),
		g.Group(children),
	)
}