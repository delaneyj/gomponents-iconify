package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeftUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.596 18.01L8.403 9.818v6.364h-2V6.404h9.778v2H9.819l8.192 8.192l-1.414 1.414Z"/>`),
		g.Group(children),
	)
}