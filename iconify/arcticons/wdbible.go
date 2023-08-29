package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wdbible(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 17.9c-5.44 7.36-12.9 14.71-22.09 14.71c-6.53 0-12.24-4.11-16.91-8.61c4.67-4.5 10.38-8.61 16.91-8.61c9.19 0 16.65 7.35 22.09 14.71"/>`),
		g.Group(children),
	)
}