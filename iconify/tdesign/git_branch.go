package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitBranch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.5 19.5a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3ZM2 18a3.5 3.5 0 1 0 4.667-3.3A2 2 0 0 1 8.5 13.5h7a4 4 0 0 0 4-4v-.145a3.502 3.502 0 0 0-1-6.855a3.5 3.5 0 0 0-1 6.855V9.5a2 2 0 0 1-2 2h-7c-.729 0-1.412.195-2 .535v-2.68a3.502 3.502 0 0 0-1-6.855a3.5 3.5 0 0 0-1 6.855v5.29A3.502 3.502 0 0 0 2 18ZM18.5 7.5a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3Zm-13 0a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3Z"/>`),
		g.Group(children),
	)
}