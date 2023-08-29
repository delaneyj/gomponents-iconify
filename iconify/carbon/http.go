package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Http(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 11h-5v10h2v-3h3a2.003 2.003 0 0 0 2-2v-3a2.002 2.002 0 0 0-2-2zm-3 5v-3h3l.001 3zm-17-3h2v8h2v-8h2v-2h-6v2zm13-2h-6v2h2v8h2v-8h2v-2zM6 11v4H3v-4H1v10h2v-4h3v4h2V11H6z"/>`),
		g.Group(children),
	)
}