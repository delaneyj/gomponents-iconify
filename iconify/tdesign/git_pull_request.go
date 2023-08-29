package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitPullRequest(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.414 3l-2 2H19.5v9.645a3.501 3.501 0 0 1-1 6.855a3.5 3.5 0 0 1-1-6.855V7h-3.086l2 2L15 10.414L10.586 6L15 1.586L16.414 3ZM5.5 4.5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3ZM2 6a3.5 3.5 0 1 1 4.5 3.355v5.29a3.501 3.501 0 0 1-1 6.855a3.5 3.5 0 0 1-1-6.855v-5.29A3.502 3.502 0 0 1 2 6Zm3.5 10.5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Zm13 0a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Z"/>`),
		g.Group(children),
	)
}