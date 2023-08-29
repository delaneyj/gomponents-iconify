package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BubbleChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTBubbleChart0"><g fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="34" cy="14" r="9"/><circle cx="12" cy="25" r="7"/><circle cx="29" cy="37" r="5"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTBubbleChart0)"/>`),
		g.Group(children),
	)
}