package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlackMediumSmallSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#354a54" d="M48.566 44.991a3.572 3.572 0 0 1-3.572 3.575H19.005c-1.972 0-3.57-1.6-3.57-3.575V19.005a3.571 3.571 0 0 1 3.57-3.572h25.989c1.975 0 3.572 1.6 3.572 3.572v25.986"/>`),
		g.Group(children),
	)
}