package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTextLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M6 22a3 3 0 0 1-3-3V5a3 3 0 0 1 3-3h8.172a3 3 0 0 1 2.12.879l3.83 3.828A3 3 0 0 1 21 8.828V19a3 3 0 0 1-3 3H6zm-1-3a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-9h-3a3 3 0 0 1-3-3V4H6a1 1 0 0 0-1 1v14zM16 8h2.586L15 4.414V7a1 1 0 0 0 1 1zm-8 4a1 1 0 1 0 0 2h8a1 1 0 1 0 0-2H8zm0 4a1 1 0 1 0 0 2h4a1 1 0 1 0 0-2H8z"/></g>`),
		g.Group(children),
	)
}