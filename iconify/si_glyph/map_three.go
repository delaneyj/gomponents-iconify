package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m11.035 13.581l4.962 1.294V3.415l-4.962-1.353v11.519zm-5.014 1.384l3.922-1.384V2.062L6.021 3.879v11.086zM.042 13.118l4.95 1.847V3.879L.042 2.031v11.087z"/>`),
		g.Group(children),
	)
}