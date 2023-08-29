package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SyncOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaSyncOutline0"><g id="evaSyncOutline1"><path id="evaSyncOutline2" fill="currentColor" d="M21.66 10.37a.62.62 0 0 0 .07-.19l.75-4a1 1 0 0 0-2-.36l-.37 2a9.22 9.22 0 0 0-16.58.84a1 1 0 0 0 .55 1.3a1 1 0 0 0 1.31-.55A7.08 7.08 0 0 1 12.07 5a7.17 7.17 0 0 1 6.24 3.58l-1.65-.27a1 1 0 1 0-.32 2l4.25.71h.16a.93.93 0 0 0 .34-.06a.33.33 0 0 0 .1-.06a.78.78 0 0 0 .2-.11l.08-.1a1.07 1.07 0 0 0 .14-.16a.58.58 0 0 0 .05-.16Zm-1.78 3.7a1 1 0 0 0-1.31.56A7.08 7.08 0 0 1 11.93 19a7.17 7.17 0 0 1-6.24-3.58l1.65.27h.16a1 1 0 0 0 .16-2L3.41 13a.91.91 0 0 0-.33 0H3a1.15 1.15 0 0 0-.32.14a1 1 0 0 0-.18.18l-.09.1a.84.84 0 0 0-.07.19a.44.44 0 0 0-.07.17l-.75 4a1 1 0 0 0 .8 1.22h.18a1 1 0 0 0 1-.82l.37-2a9.22 9.22 0 0 0 16.58-.83a1 1 0 0 0-.57-1.28Z"/></g></g>`),
		g.Group(children),
	)
}