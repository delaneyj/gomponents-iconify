package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lanxchange(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 24L5.5 5.5s8.75 17.633 9.25 18.5c-.5.868-9.25 18.5-9.25 18.5L24 24l18.5 18.5S33.75 24.868 33.25 24c.499-.864 9.25-18.5 9.25-18.5Z"/>`),
		g.Group(children),
	)
}