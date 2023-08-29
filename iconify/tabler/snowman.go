package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Snowman(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3a4 4 0 0 1 2.906 6.75a6 6 0 1 1-5.81 0A4 4 0 0 1 12 3zm5.5 8.5L20 10M6.5 11.5L4 10m8 3h.01M12 16h.01"/>`),
		g.Group(children),
	)
}