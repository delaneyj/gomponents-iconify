package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpPhotoSizeSelectActual(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M23 3H1v18h22V3zM5 17l3.5-4.5l2.5 3.01L14.5 11l4.5 6H5z"/>`),
		g.Group(children),
	)
}