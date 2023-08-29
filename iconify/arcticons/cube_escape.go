package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CubeEscape(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.86 14.3L23.89 4.5L7.03 14.23L24 24.03l16.86-9.73z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 24.03l-16.97-9.8V33.7L24 43.5V24.03zm16.86-9.73L24 24.03V43.5l16.86-9.73V14.3z"/>`),
		g.Group(children),
	)
}