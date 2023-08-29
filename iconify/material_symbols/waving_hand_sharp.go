package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WavingHandSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M.975 7q0-2.5 1.763-4.263T7 .975V3Q5.35 3 4.175 4.175T3 7H.975ZM5.3 18.725Q3.025 16.45 3.025 13.25T5.3 7.775L7.05 6l.3.3q.725.725.725 1.762T7.35 9.825l-1.05 1.05L7.9 12.5q.65.65.65 1.575T7.9 15.65l1.075 1.075q1.1-1.1 1.1-2.638T8.95 11.426l-.55-.55q.65-.65.925-1.463T9.55 7.75l5.2-5.2l1.425 1.425l-5.4 5.4l1.05 1.05l6.725-6.7l1.4 1.4l-6.7 6.725l1.05 1.05l6.025-6.025L21.75 8.3l-6.025 6.025l1.05 1.05L21.55 10.6l1.425 1.425l-6.725 6.7Q13.975 21 10.775 21T5.3 18.725Zm11.7 4.3V21q1.65 0 2.825-1.175T21 17h2.025q0 2.5-1.763 4.263T17 23.025Z"/>`),
		g.Group(children),
	)
}