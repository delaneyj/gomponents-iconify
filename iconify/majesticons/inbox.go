package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Inbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M6 21a3 3 0 0 1-3-3v-3h5a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2h5v3a3 3 0 0 1-3 3H6zm15-8h-5a2 2 0 0 0-2 2h-4a2 2 0 0 0-2-2H3V6a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v7z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}