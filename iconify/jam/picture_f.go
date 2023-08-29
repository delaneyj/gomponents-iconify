package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PictureF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20 10.536l-4.416-4.44a3 3 0 0 0-4.69.582L5.072 16H3a3 3 0 0 1-3-3V3a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v7.536zm-.011 2.724A3 3 0 0 1 17 16H7.64l4.969-8.293a1 1 0 0 1 1.563-.195l5.817 5.748zM6 9a3 3 0 1 0 0-6a3 3 0 0 0 0 6z"/>`),
		g.Group(children),
	)
}