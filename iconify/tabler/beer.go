package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Beer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 21h6a1 1 0 0 0 1-1v-3.625c0-1.397.29-2.775.845-4.025l.31-.7C17.711 10.4 18 9.397 18 8V4a1 1 0 0 0-1-1H7a1 1 0 0 0-1 1v4c0 1.397.29 2.4.845 3.65l.31.7A9.931 9.931 0 0 1 8 16.375V20a1 1 0 0 0 1 1zM6 8h12"/>`),
		g.Group(children),
	)
}