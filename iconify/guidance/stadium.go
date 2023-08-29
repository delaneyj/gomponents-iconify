package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stadium(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M12 3.5c-2.663 0-5.323-.189-7.96-.566L1 2.5H.5v19H1l3.04-.434A56.277 56.277 0 0 1 12 20.5m0-17c2.663 0 5.323-.189 7.96-.566L23 2.5h.5v19H23l-3.04-.434A56.276 56.276 0 0 0 12 20.5m0-17V9m0 11.5V15M.5 7.5H1a4.5 4.5 0 0 1 0 9H.5m23 0H23a4.5 4.5 0 1 1 0-9h.5M12 9a3 3 0 1 1 0 6m0-6a3 3 0 1 0 0 6"/>`),
		g.Group(children),
	)
}