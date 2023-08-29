package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PetriDish(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M22 14.5a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0ZM11.5 13a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm4.5 6.5a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0Z"/><path d="m4 15l.008-.03A13.242 13.242 0 0 1 4 14.5C4 7.596 9.373 2 16 2s12 5.596 12 12.5c0 .16-.003.318-.009.476L28 15v3c0 5.799-4 12-12 12S4 23.799 4 18c0 0 .022-3.08 0-3Zm23 1c0-5.523-5.189-10-11-10S5 10.477 5 16s5.189 10 11 10s11-4.477 11-10Z"/></g>`),
		g.Group(children),
	)
}