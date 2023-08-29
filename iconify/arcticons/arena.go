package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Arena(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 31.07c-3.44-4.52-9.75-9-9.75-9"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.18 41.2A76.38 76.38 0 0 0 43.5 30.52S28.72 20.58 24 6.8C19.28 20.58 4.5 30.52 4.5 30.52A76.38 76.38 0 0 0 17.82 41.2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 31.07c3.44-4.52 9.75-9 9.75-9"/>`),
		g.Group(children),
	)
}