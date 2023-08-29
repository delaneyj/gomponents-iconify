package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConnectionArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M40.99 8.027H12.182C9.455 8.027 4 9.55 4 15.967C4 22.384 9.455 24 12.182 24h23.812C38.72 24 44 25.568 44 31.985c0 6.416-5.279 8.018-8.006 8.018H6.065"/><path d="m8.046 35.993l-3.979 4.066L8.046 44M38.043 3.954L42.02 8.02l-3.978 3.941"/></g>`),
		g.Group(children),
	)
}