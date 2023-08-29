package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicNotesPlusFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M232 56a8 8 0 0 1-8 8h-16v16a8 8 0 0 1-16 0V64h-16a8 8 0 0 1 0-16h16V32a8 8 0 0 1 16 0v16h16a8 8 0 0 1 8 8Zm-24 56a8 8 0 0 0-8 8v22.08A36 36 0 1 0 216 172v-52a8 8 0 0 0-8-8Zm-54.42-10.67a8 8 0 0 0 2.76-9.88a8.11 8.11 0 0 0-1.1-1.79a55.78 55.78 0 0 1-11-39A8 8 0 0 0 137 42a7.9 7.9 0 0 0-2.61.21L78.06 56.24A8 8 0 0 0 72 64v110.08A36 36 0 1 0 88 204v-85.75l62.82-15.71a8.06 8.06 0 0 0 2.76-1.21Z"/>`),
		g.Group(children),
	)
}