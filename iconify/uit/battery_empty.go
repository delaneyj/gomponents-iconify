package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryEmpty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.5 10a.5.5 0 0 0-.5.5v3a.5.5 0 1 0 1 0v-3a.5.5 0 0 0-.5-.5zm-4-3h-14A1.5 1.5 0 0 0 2 8.5v7A1.5 1.5 0 0 0 3.5 17h14a1.5 1.5 0 0 0 1.5-1.5v-7A1.5 1.5 0 0 0 17.5 7zm.5 8.5a.5.5 0 0 1-.5.5h-14a.501.501 0 0 1-.5-.5v-7a.5.5 0 0 1 .5-.5h14a.5.5 0 0 1 .5.5v7z"/>`),
		g.Group(children),
	)
}