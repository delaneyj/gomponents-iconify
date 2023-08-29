package nimbus

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sun(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 3.33A4.85 4.85 0 0 0 3 8a4.85 4.85 0 0 0 5 4.67A4.85 4.85 0 0 0 13 8a4.85 4.85 0 0 0-5-4.67zm0 8.09A3.6 3.6 0 0 1 4.25 8A3.6 3.6 0 0 1 8 4.58A3.6 3.6 0 0 1 11.75 8A3.6 3.6 0 0 1 8 11.42zM7.33 0h1.25v2H7.33zm0 14h1.25v2H7.33zM14 7.33h2v1.25h-2zm-14 0h2v1.25H0zM11.92 2.4h2v1.25h-2zm-9.9 9.9h2v1.25h-2zm10.33-.38h1.25v2h-1.25zm-9.9-9.9H3.7v2H2.45z"/>`),
		g.Group(children),
	)
}