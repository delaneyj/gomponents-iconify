package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pzizz(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 32.535c12.601.088 19.573-8.015 19.5-17.07h-7.686C34.785 17.934 33.092 23.86 24 24c-9.092-.14-10.785-6.065-11.814-8.536H4.501C4.427 24.52 11.398 32.623 24 32.535Z"/>`),
		g.Group(children),
	)
}