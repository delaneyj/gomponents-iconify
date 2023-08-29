package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MuseumTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6 1.798l4 2.667v6.355l12-2V22h-6v-4h-2v4H2V4.465l4-2.667ZM4 5.535V20h8v-4h6v4h2v-8.82l-12 2V5.536L6 4.202L4 5.535Z"/>`),
		g.Group(children),
	)
}