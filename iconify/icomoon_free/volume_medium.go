package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeMedium(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M11.243 12.993a.75.75 0 0 1-.53-1.281a5.256 5.256 0 0 0 0-7.425a.75.75 0 1 1 1.061-1.061c1.275 1.275 1.977 2.97 1.977 4.773s-.702 3.498-1.977 4.773a.748.748 0 0 1-.53.22zm-2.665-1.415a.75.75 0 0 1-.53-1.281a3.254 3.254 0 0 0 0-4.596a.75.75 0 1 1 1.061-1.061a4.756 4.756 0 0 1 0 6.718a.748.748 0 0 1-.53.22zM6.5 15a.504.504 0 0 1-.354-.146L2.292 11H.499a.5.5 0 0 1-.5-.5v-5a.5.5 0 0 1 .5-.5h1.793l3.854-3.854A.499.499 0 0 1 7 1.5v13a.5.5 0 0 1-.5.5z"/>`),
		g.Group(children),
	)
}