package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowThinUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m9.002.02l6.693 6.305c.459.459.297 1.16.094 1.359c-.402.402-.998.342-1.402-.061L9.949 3.728l-.002 11.314a.941.941 0 0 1-.951.928a.94.94 0 0 1-.951-.928L8.033 3.831L3.762 7.685a1.036 1.036 0 0 1-1.461 0c-.404-.4-.303-.949.102-1.352L9.002.02z"/>`),
		g.Group(children),
	)
}