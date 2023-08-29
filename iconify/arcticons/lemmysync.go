package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lemmysync(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.47 42.5L5.91 25.62h25.41L18.47 42.5zm-1.73-19.92L29.22 5.5l12.87 17.08H16.74z"/>`),
		g.Group(children),
	)
}