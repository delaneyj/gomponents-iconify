package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookOpenThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16 26.329A4.491 4.491 0 0 1 12.5 28H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h8.5c1.414 0 2.675.652 3.5 1.671A4.491 4.491 0 0 1 19.5 4H28a2 2 0 0 1 2 2v20a2 2 0 0 1-2 2h-8.5a4.491 4.491 0 0 1-3.5-1.671ZM17 8.5v15a2.5 2.5 0 0 0 2.5 2.5H28V6h-8.5A2.5 2.5 0 0 0 17 8.5Zm-2 0A2.5 2.5 0 0 0 12.5 6H4v20h8.5a2.5 2.5 0 0 0 2.5-2.5v-15Z"/>`),
		g.Group(children),
	)
}