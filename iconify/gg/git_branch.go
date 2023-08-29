package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitBranch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 3a2 2 0 0 0-1 3.732v10.536A2 2 0 1 0 10.732 20h1.227a4 4 0 0 0 4-4v-1.268a2 2 0 0 0-1-3.732a2 2 0 0 0-1 3.732V16a2 2 0 0 1-2 2h-1.227a2.01 2.01 0 0 0-.732-.732V6.732A2 2 0 0 0 9 3Z"/>`),
		g.Group(children),
	)
}