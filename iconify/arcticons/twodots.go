package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Twodots(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="34.5" cy="13.5" r="8" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="13.5" cy="34.5" r="8" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}