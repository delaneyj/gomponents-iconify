package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InboxIn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M9.707 7.293a1 1 0 0 0-1.414 1.414l3 3a1 1 0 0 0 1.414 0l3-3a1 1 0 0 0-1.414-1.414L13 8.586V3h5a3 3 0 0 1 3 3v7h-5a2 2 0 0 0-2 2h-4a2 2 0 0 0-2-2H3V6a3 3 0 0 1 3-3h5v5.586L9.707 7.293zM3 18a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-3h-5a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2H3v3z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}