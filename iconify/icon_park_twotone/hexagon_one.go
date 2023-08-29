package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HexagonOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTHexagonOne0"><path fill="#555" stroke="#fff" stroke-width="4" d="m23.029 43.46l-16-8.889A2 2 0 0 1 6 32.824V15.177a2 2 0 0 1 1.029-1.748l16-8.89a2 2 0 0 1 1.942 0l16 8.89A2 2 0 0 1 42 15.176v17.646a2 2 0 0 1-1.029 1.748l-16 8.89a2 2 0 0 1-1.942 0Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTHexagonOne0)"/>`),
		g.Group(children),
	)
}