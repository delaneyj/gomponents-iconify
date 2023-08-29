package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareCircleBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M12 3a3 3 0 1 0 3 3m-9.5 9a3 3 0 1 0 3 3m10 3a3 3 0 1 1 3-3M20 13a7.98 7.98 0 0 0-2.708-6M4 13a7.98 7.98 0 0 1 2.708-6M10 20.748A8 8 0 0 0 12 21a8 8 0 0 0 2-.252"/>`),
		g.Group(children),
	)
}