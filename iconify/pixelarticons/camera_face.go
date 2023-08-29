package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 3h10v2h5v16H2V7h2v12h16V7h-5V5H9v2H2V5h5V3zm7 12h-4v2h4v-2zm-4-2v2H8v-2h2zm0-2V9H8v2h2zm6 2v2h-2v-2h2zm0-2V9h-2v2h2z"/>`),
		g.Group(children),
	)
}