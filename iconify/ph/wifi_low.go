package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiLow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M140 204a12 12 0 1 1-12-12a12 12 0 0 1 12 12Zm32.71-45.47a76.05 76.05 0 0 0-89.42 0a8 8 0 0 0 9.42 12.94a60 60 0 0 1 70.58 0a8 8 0 1 0 9.42-12.94Z"/>`),
		g.Group(children),
	)
}