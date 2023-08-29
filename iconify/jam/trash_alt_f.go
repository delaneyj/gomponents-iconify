package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrashAltF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14.778 6l-.673 12.11A2 2 0 0 1 12.108 20H3.892a2 2 0 0 1-1.997-1.89L1.222 6h13.556zM1 0h14a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1H1a1 1 0 0 1-1-1V1a1 1 0 0 1 1-1zm5 8a1 1 0 0 0-1 1v7a1 1 0 0 0 2 0V9a1 1 0 0 0-1-1zm4 0a1 1 0 0 0-1 1v7a1 1 0 0 0 2 0V9a1 1 0 0 0-1-1z"/>`),
		g.Group(children),
	)
}