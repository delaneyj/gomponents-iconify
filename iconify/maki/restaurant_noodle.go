package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RestaurantNoodle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M4.457 11.99L1 8V7h13v1l-3.496 3.99ZM3.988 2.5a.5.5 0 0 0-1 0v.567l-1.797.368a.25.25 0 1 0 .094.49l1.703-.277v.566l-1.75.036a.25.25 0 0 0 0 .5l1.75.036v1.212h1Zm9.5 1.5l-7.5.263V2.995l7.594-1.074a.5.5 0 0 0-.188-.982L5.98 2.455a.496.496 0 0 0-.99.045v.228l-.494.1v.352l.493-.08v1.197l-.493.01v.461L13.488 5a.5.5 0 0 0 0-1ZM10 13H5v.576h5Z"/>`),
		g.Group(children),
	)
}