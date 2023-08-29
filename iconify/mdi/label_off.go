package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LabelOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 4.27L3.28 3L20 19.72L18.73 21l-2.1-2.1c-.2.06-.41.1-.63.1H5a2 2 0 0 1-2-2V7c0-.5.17-.93.46-1.27L2 4.27m15.63 1.57L22 12l-3 4.2L7.83 5H16c.67 0 1.27.33 1.63.84Z"/>`),
		g.Group(children),
	)
}