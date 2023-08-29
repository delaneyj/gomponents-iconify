package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitCommit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11 4a1 1 0 1 1 2 0v5.17a3.001 3.001 0 0 1 0 5.66V20a1 1 0 1 1-2 0v-5.17a3.001 3.001 0 0 1 0-5.66V4Zm1 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}