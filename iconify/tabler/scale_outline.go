package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScaleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 7a4 4 0 0 1 4-4h10a4 4 0 0 1 4 4v10a4 4 0 0 1-4 4H7a4 4 0 0 1-4-4z"/><path d="M12 7c1.956 0 3.724.802 5 2.095l-2.956 2.904a3 3 0 0 0-2.038-.799a3 3 0 0 0-2.038.798L7.012 9.095a6.979 6.979 0 0 1 5-2.095z"/></g>`),
		g.Group(children),
	)
}