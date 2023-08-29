package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitMerge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.5 4.5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3ZM2 6a3.5 3.5 0 1 1 4.667 3.3A2 2 0 0 0 8.5 10.5h7a4 4 0 0 1 4 4v.145a3.502 3.502 0 0 1-1 6.855a3.5 3.5 0 0 1-1-6.855V14.5a2 2 0 0 0-2-2h-7a3.982 3.982 0 0 1-2-.535v2.68a3.502 3.502 0 0 1-1 6.855a3.5 3.5 0 0 1-1-6.855v-5.29A3.502 3.502 0 0 1 2 6Zm16.5 10.5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Zm-13 0a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Z"/>`),
		g.Group(children),
	)
}