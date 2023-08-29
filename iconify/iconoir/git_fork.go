package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitFork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 7a2 2 0 1 0 0-4a2 2 0 0 0 0 4ZM7 7a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm0 14a2 2 0 1 0 0-4a2 2 0 0 0 0 4ZM7 7v10M17 7v1c0 2.5-2 3-2 3l-6 2s-2 .5-2 3v1"/>`),
		g.Group(children),
	)
}