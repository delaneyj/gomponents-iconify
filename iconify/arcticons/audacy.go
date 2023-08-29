package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Audacy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="13.619" cy="30.071" r="9.119" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.668 39.27H43.5L27.152 8.73H14.08c5.781 10.018 10.71 20.355 16.588 30.54Z"/>`),
		g.Group(children),
	)
}