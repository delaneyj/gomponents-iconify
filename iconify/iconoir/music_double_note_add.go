package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicDoubleNoteAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 6.5h3m3 0h-3m0 0v-3m0 3v3M6 16V5l8-1m1 10v-4m-3 9h1a2 2 0 0 0 2-2v-3h-3a2 2 0 0 0-2 2v1a2 2 0 0 0 2 2Zm-9 2h1a2 2 0 0 0 2-2v-3H3a2 2 0 0 0-2 2v1a2 2 0 0 0 2 2Z"/>`),
		g.Group(children),
	)
}