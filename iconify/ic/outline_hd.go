package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OutlineHd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 3H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14c1.1 0 2-.9 2-2V5c0-1.1-.9-2-2-2zm0 16H5V5h14v14zM7.5 13h2v2H11V9H9.5v2.5h-2V9H6v6h1.5zM18 14v-4c0-.55-.45-1-1-1h-4v6h4c.55 0 1-.45 1-1zm-1.5-.5h-2v-3h2v3z"/>`),
		g.Group(children),
	)
}