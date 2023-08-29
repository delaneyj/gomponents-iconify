package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Green(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M24.485 2c0 8-18 4-18 20c0 6 2 8 2 8h2s-3-2-3-8c0-4 9-8 9-8s-7.98 4.328-7.98 8.436C21.238 24.43 28.287 9.606 24.484 2z"/>`),
		g.Group(children),
	)
}