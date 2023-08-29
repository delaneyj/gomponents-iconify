package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicNote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M9 18c0 1.105-1.12 2-2.5 2S4 19.105 4 18s1.12-2 2.5-2s2.5.895 2.5 2zm0 0V7l11-3v11m0 0c0 1.105-1.12 2-2.5 2s-2.5-.895-2.5-2s1.12-2 2.5-2s2.5.895 2.5 2z"/>`),
		g.Group(children),
	)
}