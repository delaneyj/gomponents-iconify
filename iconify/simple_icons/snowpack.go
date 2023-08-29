package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Snowpack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m23.81 19.751l-10.8-16.8a1.2 1.2 0 0 0-2.02 0l-10.8 16.8a1.2 1.2 0 0 0-.043 1.224A1.2 1.2 0 0 0 1.2 21.6h21.6a1.2 1.2 0 0 0 1.053-.625a1.2 1.2 0 0 0-.044-1.224zM12 5.82l3.202 4.981H12l-2.4 2.4l-1.427-1.427z"/>`),
		g.Group(children),
	)
}