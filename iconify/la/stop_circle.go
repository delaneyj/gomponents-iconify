package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StopCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 3C8.85 3 3 8.85 3 16s5.85 13 13 13s13-5.85 13-13S23.15 3 16 3zm0 2c6.05 0 11 4.95 11 11s-4.95 11-11 11S5 22.05 5 16S9.95 5 16 5zm-5 6v10h10V11H11zm2 2h6v6h-6v-6z"/>`),
		g.Group(children),
	)
}