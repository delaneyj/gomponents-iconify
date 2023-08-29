package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitMergeOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M2.5 10.5a2 2 0 1 0 0 4a2 2 0 0 0 0-4Zm0 0v-6m2-2a2 2 0 1 0-2 2m2-2a2 2 0 0 1-2 2m2-2h5a3 3 0 0 1 3 3v2m0 0a2 2 0 1 0 0 4a2 2 0 0 0 0-4Z"/>`),
		g.Group(children),
	)
}