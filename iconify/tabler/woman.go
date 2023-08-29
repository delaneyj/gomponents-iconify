package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Woman(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 16v5m4-5v5m-6-5h8l-2-7h-4zm-3-5c1.667-1.333 3.333-2 5-2m9 2c-1.667-1.333-3.333-2-5-2m-4-5a2 2 0 1 0 4 0a2 2 0 1 0-4 0"/>`),
		g.Group(children),
	)
}