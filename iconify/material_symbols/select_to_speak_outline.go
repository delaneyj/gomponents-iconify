package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SelectToSpeakOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 20.725v-2.05q2.25-.65 3.625-2.5t1.375-4.2q0-2.35-1.375-4.2T14 5.275v-2.05q3.1.7 5.05 3.138T21 11.975q0 3.175-1.95 5.613T14 20.725ZM12 20l-5-5H3V9h4l5-5v16Zm2-4V7.95q1.175.55 1.838 1.65T16.5 12q0 1.275-.663 2.363T14 16Zm-4-.85v-6.3L7.85 11H5v2h2.85L10 15.15ZM1 7V1h6v2H3v4H1Zm16 16v-2h4v-4h2v6h-6ZM7.5 12Z"/>`),
		g.Group(children),
	)
}