package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanelLeftHeaderTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5 3a3 3 0 0 0-3 3v7a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3H5Zm12 3v1H8.5V4H15a2 2 0 0 1 2 2Zm-2 9H8.5V8H17v5a2 2 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}