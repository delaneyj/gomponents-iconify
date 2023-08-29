package svg_spinners

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThreeDotsScaleMiddle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<circle cx="4" cy="12" r="1.5" fill="currentColor"><animate attributeName="r" dur="0.75s" repeatCount="indefinite" values="1.5;3;1.5"/></circle><circle cx="12" cy="12" r="3" fill="currentColor"><animate attributeName="r" dur="0.75s" repeatCount="indefinite" values="3;1.5;3"/></circle><circle cx="20" cy="12" r="1.5" fill="currentColor"><animate attributeName="r" dur="0.75s" repeatCount="indefinite" values="1.5;3;1.5"/></circle>`),
		g.Group(children),
	)
}