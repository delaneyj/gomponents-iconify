package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Coffee(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 15c0 1.1.9 2 2 2h10c1.1 0 2-.9 2-2H4zm2-3.8v.8c0 1.1.9 2 2 2h7c1.1 0 2-.9 2-2V4H5.6C3.6 4 2 5.4 2 7.6c0 2.3 1.6 3.6 3.6 3.6H6zM3.9 7.6c0-1 .7-1.8 1.7-1.8H6v3.6h-.4c-1 0-1.7-.7-1.7-1.8z"/>`),
		g.Group(children),
	)
}