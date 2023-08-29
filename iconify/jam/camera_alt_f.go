package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraAltF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 12H0V6a6 6 0 0 1 6-6h8a6 6 0 0 1 6 6v6zm0 2a6 6 0 0 1-6 6H6a6 6 0 0 1-6-6h20zM6 15a1 1 0 0 0 0 2h8a1 1 0 0 0 0-2H6zm4-4a4 4 0 1 0 0-8a4 4 0 0 0 0 8zm6 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2zm-6-2a2 2 0 1 1 0-4a2 2 0 0 1 0 4z"/>`),
		g.Group(children),
	)
}