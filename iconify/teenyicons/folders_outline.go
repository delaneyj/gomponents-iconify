package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoldersOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M3.5 8.5v-7a1 1 0 0 1 1-1h3l2 2h4a1 1 0 0 1 1 1v5a1 1 0 0 1-1 1h-9m-1-1a1 1 0 0 0 1 1m-1-1v-3h-2a1 1 0 0 0-1 1v7a1 1 0 0 0 1 1h9a1 1 0 0 0 1-1v-4h-7"/>`),
		g.Group(children),
	)
}