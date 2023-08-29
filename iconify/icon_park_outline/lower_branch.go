package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LowerBranch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M6 23c0-1.5 1.999-5.5 7.061-6.571C18.18 15.346 22.848 10.857 24 9m18 14c.012-1.5-2-5.5-7.062-6.571C29.821 15.346 25.152 10.857 24 9"/><circle r="4" fill="currentColor" transform="matrix(0 1 1 0 24 9)"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 9v14M5 33h6m-6 8h6m26-8h6m-6 8h6m-22-8h6m-6 8h6"/></g>`),
		g.Group(children),
	)
}