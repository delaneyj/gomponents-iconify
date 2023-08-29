package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dnshero(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.3 16.15L4.5 35.85l35.6 3.9l-17.6-7.4L43.5 11L7.3 8.25Z"/>`),
		g.Group(children),
	)
}