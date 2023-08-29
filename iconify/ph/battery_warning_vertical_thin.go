package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryWarningVerticalThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M124 136V96a4 4 0 0 1 8 0v40a4 4 0 0 1-8 0Zm4 28a8 8 0 1 0 8 8a8 8 0 0 0-8-8ZM96 12h64a4 4 0 0 0 0-8H96a4 4 0 0 0 0 8Zm100 44v168a20 20 0 0 1-20 20H80a20 20 0 0 1-20-20V56a20 20 0 0 1 20-20h96a20 20 0 0 1 20 20Zm-8 0a12 12 0 0 0-12-12H80a12 12 0 0 0-12 12v168a12 12 0 0 0 12 12h96a12 12 0 0 0 12-12Z"/>`),
		g.Group(children),
	)
}