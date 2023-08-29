package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RuleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" d="M9 16h14v2H9zm0-6h14v2H9z"/><path fill="currentColor" d="M26 2H6a2.002 2.002 0 0 0-2 2v13a10.981 10.981 0 0 0 5.824 9.707L16 30l6.176-3.293A10.981 10.981 0 0 0 28 17V4a2.002 2.002 0 0 0-2-2Zm-3 16H9v-2h14Zm0-6H9v-2h14Z"/>`),
		g.Group(children),
	)
}