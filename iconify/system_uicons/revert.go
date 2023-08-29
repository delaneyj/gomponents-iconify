package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Revert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m7.5 5.5l-3-3h6a8.003 8.003 0 0 1 7.427 5.02c.37.921.573 1.927.573 2.98a8 8 0 1 1-16 0c0-1.49.37-3.472 1.538-5.091"/><path d="M10.5 5.5v5h3"/></g>`),
		g.Group(children),
	)
}