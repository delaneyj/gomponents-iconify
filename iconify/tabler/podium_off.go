package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PodiumOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8h7l-.621 2.485A2 2 0 0 1 16.439 12H16m-4 0H7.561a2 2 0 0 1-1.94-1.515L5 8h3M7 8V7m.864-3.106A2.99 2.99 0 0 1 10 3m-2 9l1 9m6.599-5.387L15 21m-8 0h10M3 3l18 18"/>`),
		g.Group(children),
	)
}