package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LemonTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 4a2 2 0 0 1 1.185 3.611c1.55 2.94.873 6.917-1.892 9.682c-2.765 2.765-6.743 3.442-9.682 1.892a2 2 0 1 1-2.796-2.796c-1.55-2.94-.873-6.917 1.892-9.682c2.765-2.765 6.743-3.442 9.682-1.892A2 2 0 0 1 18 4z"/>`),
		g.Group(children),
	)
}