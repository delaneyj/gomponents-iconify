package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaddingTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 0v16h16V0H0zm15 3h-1v1h1v11H1V3h1V2H1V1h1v1h1V1h1v1h1V1h1v1h1V1h1v1h1V1h1v1h1V1h1v1h1V1h1v1h1v1z"/><path fill="currentColor" d="M3 2h1v1H3V2zM2 3h1v1H2V3zm2 0h1v1H4V3zm2 0h1v1H6V3zM5 2h1v1H5V2zm2 0h1v1H7V2zm2 0h1v1H9V2zM8 3h1v1H8V3zm2 0h1v1h-1V3zm2 0h1v1h-1V3zm-1-1h1v1h-1V2zm2 0h1v1h-1V2z"/>`),
		g.Group(children),
	)
}