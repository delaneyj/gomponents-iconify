package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MobileRetro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M11 0h-1v2H4v14h7V0zM6 14H5v-1h1v1zm0-2H5v-1h1v1zm0-2H5V9h1v1zm2 4H7v-1h1v1zm0-2H7v-1h1v1zm0-2H7V9h1v1zm2 4H9v-1h1v1zm0-2H9v-1h1v1zm0-2H9V9h1v1zm0-2H5V4h5v4z"/>`),
		g.Group(children),
	)
}