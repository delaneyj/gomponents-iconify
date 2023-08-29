package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicNotesFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M212.92 25.71a7.89 7.89 0 0 0-6.86-1.46l-128 32A8 8 0 0 0 72 64v110.1A36 36 0 1 0 88 204v-93.75l112-28v59.85a36 36 0 1 0 16 29.9V32a8 8 0 0 0-3.08-6.29Z"/>`),
		g.Group(children),
	)
}