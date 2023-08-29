package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitPullOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="m8.5.5l-2 2m0 0l2 2m-2-2h3a3 3 0 0 1 3 3v2m-10 3a2 2 0 1 0 0 4a2 2 0 0 0 0-4Zm0 0v-6m0 0a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm10 3a2 2 0 1 0 0 4a2 2 0 0 0 0-4Z"/>`),
		g.Group(children),
	)
}