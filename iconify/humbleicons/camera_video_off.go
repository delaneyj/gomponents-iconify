package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraVideoOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 5l14 14M9 7h6a1 1 0 0 1 1 1v5.5a1 1 0 0 0 .4.8L20 17V7l-4 3m-1 7H5a1 1 0 0 1-1-1V8a1 1 0 0 1 1-1"/>`),
		g.Group(children),
	)
}