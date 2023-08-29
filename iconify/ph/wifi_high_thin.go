package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiHighThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M136 204a8 8 0 1 1-8-8a8 8 0 0 1 8 8Zm98.54-113.9a168 168 0 0 0-213.08 0a4 4 0 1 0 5.08 6.18a160 160 0 0 1 202.92 0a4 4 0 0 0 5.08-6.18Zm-32.06 35.81a120 120 0 0 0-149 0a4 4 0 0 0 5 6.27a112 112 0 0 1 139 0a4 4 0 0 0 5-6.27Zm-32.13 35.86a72 72 0 0 0-84.7 0a4 4 0 1 0 4.7 6.46a64.07 64.07 0 0 1 75.3 0a4 4 0 0 0 5.58-.87a4 4 0 0 0-.88-5.59Z"/>`),
		g.Group(children),
	)
}