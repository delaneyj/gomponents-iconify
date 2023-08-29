package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CellularDataCellularOffTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M3.28 2.22a.75.75 0 1 0-1.06 1.06L11 12.06v7.19l.007.102a.75.75 0 0 0 1.493-.102v-5.69l2.5 2.5v3.19l.007.102a.75.75 0 0 0 1.493-.102v-1.69l4.22 4.22a.75.75 0 0 0 1.06-1.06L3.28 2.22z" fill="currentColor"/><path d="M19 15.818l1.5 1.5V5.742l-.008-.101A.75.75 0 0 0 19 5.757v10.061z" fill="currentColor"/><path d="M15 11.818l1.5 1.5V8.75l-.007-.102A.75.75 0 0 0 15 8.75v3.068z" fill="currentColor"/><path d="M3.75 17a.75.75 0 0 1 .743.648l.007.102v1.5a.75.75 0 0 1-1.493.102L3 19.25v-1.5a.75.75 0 0 1 .75-.75z" fill="currentColor"/><path d="M7.75 14a.75.75 0 0 1 .743.648l.007.102v4.499a.75.75 0 0 1-1.493.102L7 19.248V14.75a.75.75 0 0 1 .75-.75z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}