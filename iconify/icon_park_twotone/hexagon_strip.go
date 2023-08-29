package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HexagonStrip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTHexagonStrip0"><path fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M19 4h10v11.34l9.82-5.67l5 8.66L34 24l9.82 5.67l-5 8.66L29 32.66V44H19V32.66l-9.82 5.67l-5-8.66L14 24l-9.82-5.67l5-8.66L19 15.34V4Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTHexagonStrip0)"/>`),
		g.Group(children),
	)
}