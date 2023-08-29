package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GameControllerSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M0 5.5A3.5 3.5 0 0 1 3.5 2h8A3.5 3.5 0 0 1 15 5.5v4.528a2.972 2.972 0 0 1-5.63 1.329L9.19 11H5.809l-.179.357A2.972 2.972 0 0 1 0 10.027V5.5ZM4 8V7H3V6h1V5h1v1h1v1H5v1H4Zm6 0H9V7h1v1Zm1-2h1V5h-1v1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}