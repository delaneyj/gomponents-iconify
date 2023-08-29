package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiMediumThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M136 204a8 8 0 1 1-8-8a8 8 0 0 1 8 8Zm66.48-78.09a120 120 0 0 0-149 0a4 4 0 0 0 5 6.27a112 112 0 0 1 139 0a4 4 0 0 0 5-6.27Zm-32.13 35.86a72 72 0 0 0-84.7 0a4 4 0 1 0 4.71 6.46a64 64 0 0 1 75.28 0a4 4 0 1 0 4.71-6.46Z"/>`),
		g.Group(children),
	)
}