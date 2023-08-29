package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15 4h2.94c.59 0 1.06.47 1.06 1.06v11.81c0 .59-.47 1.13-1.06 1.13H2.06C1.47 18 1 17.46 1 16.87V5.06C1 4.47 1.47 4 2.06 4H5l3-2h4zm-5 11c2.21 0 4-1.79 4-4s-1.79-4-4-4s-4 1.79-4 4s1.79 4 4 4z"/>`),
		g.Group(children),
	)
}