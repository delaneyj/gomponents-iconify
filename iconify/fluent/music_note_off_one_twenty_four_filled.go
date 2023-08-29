package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicNoteOffOneTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m12 13.06l8.72 8.72a.75.75 0 0 0 1.06-1.06L3.28 2.22a.75.75 0 1 0-1.06 1.06l8.28 8.28v3.317a4 4 0 1 0 1.476 3.56a.749.749 0 0 0 .024-.187v-5.19Zm6.987-1.608l-6.964-2.611L10.5 7.318V2.75a.75.75 0 0 1 1.013-.702l6.378 2.391A3.25 3.25 0 0 1 20 7.483v3.267a.75.75 0 0 1-1.013.702Z"/>`),
		g.Group(children),
	)
}