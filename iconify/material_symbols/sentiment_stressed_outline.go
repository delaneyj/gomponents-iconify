package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SentimentStressedOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m7.45 12.6l3.575-2.6L7.45 7.4l-.9 1.2L8.475 10L6.55 11.4l.9 1.2Zm9.1 0l.9-1.2l-1.925-1.4l1.925-1.4l-.9-1.2l-3.575 2.6l3.575 2.6Zm-6.05 4.45l1.5-1.5l1.5 1.5l1.5-1.5l.975.975l1.05-1.05L15 13.45l-1.5 1.5l-1.5-1.5l-1.5 1.5l-1.5-1.5l-2.025 2.025l1.05 1.05L9 15.55l1.5 1.5ZM12 22q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-2.075.788-3.9t2.137-3.175q1.35-1.35 3.175-2.137T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12q0 2.075-.788 3.9t-2.137 3.175q-1.35 1.35-3.175 2.138T12 22Zm0-10Zm0 8q3.35 0 5.675-2.325T20 12q0-3.35-2.325-5.675T12 4Q8.65 4 6.325 6.325T4 12q0 3.35 2.325 5.675T12 20Z"/>`),
		g.Group(children),
	)
}