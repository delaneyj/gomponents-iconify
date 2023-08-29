package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fingerprint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 11C2 5.477 6.477 1 12 1s10 4.477 10 10v3c0 2.64-.57 5.15-1.592 7.412l-.413.911l-1.822-.824l.412-.911A15.938 15.938 0 0 0 20 14v-3a8 8 0 1 0-16 0v3H2v-3Zm4.5 0a5.5 5.5 0 1 1 11 0v3c0 2.8-.853 5.403-2.314 7.56l-.561.829l-1.656-1.122l.56-.828A11.442 11.442 0 0 0 15.5 14v-3a3.5 3.5 0 1 0-7 0v3A4.5 4.5 0 0 1 4 18.5H2v-2h2A2.5 2.5 0 0 0 6.5 14v-3Zm6.5-1v4c0 3.39-1.875 6.34-4.639 7.874l-.874.486l-.97-1.75l.874-.484A6.998 6.998 0 0 0 11 14v-4h2Z"/>`),
		g.Group(children),
	)
}