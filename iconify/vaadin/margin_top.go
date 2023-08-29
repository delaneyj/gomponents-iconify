package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarginTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M15 2v1h-1V2h-1v1h-1V2h-1v1h-1V2H9v1H8V2H7v1H6V2H5v1H4V2H3v1H2V2H1v1H0v13h16V2h-1zm0 13H1V4h14v11zm0-15h1v1h-1V0zm-1 1h1v1h-1V1zm-1-1h1v1h-1V0zm-1 1h1v1h-1V1zm-1-1h1v1h-1V0zm-1 1h1v1h-1V1zM9 0h1v1H9V0zM8 1h1v1H8V1zM7 0h1v1H7V0zM6 1h1v1H6V1zM5 0h1v1H5V0zM4 1h1v1H4V1zM3 0h1v1H3V0zM2 1h1v1H2V1zM1 0h1v1H1V0zM0 1h1v1H0V1z"/>`),
		g.Group(children),
	)
}