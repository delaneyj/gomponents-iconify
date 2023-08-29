package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicNote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.12 2.21a1 1 0 0 0-.86-.21l-8 2a1 1 0 0 0-.76 1v10.35A3.45 3.45 0 0 0 8 15a3.5 3.5 0 1 0 3.5 3.5v-7.72L18.74 9h.07l.19-.15l.15-.1a.93.93 0 0 0 .13-.15a.78.78 0 0 0 .1-.15a.55.55 0 0 0 .06-.18a.58.58 0 0 0 0-.19a.24.24 0 0 0 0-.08V3a1 1 0 0 0-.32-.79ZM8 20a1.5 1.5 0 1 1 1.5-1.5A1.5 1.5 0 0 1 8 20Zm9.5-12.78l-6 1.5V5.78l6-1.5Z"/>`),
		g.Group(children),
	)
}