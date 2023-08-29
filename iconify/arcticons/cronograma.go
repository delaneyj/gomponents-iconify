package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cronograma(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.5 14.21h25m-25 6.53h4.28m-4.28 6.52h4.28m-4.28 6.53h4.28m17.46-13.05h3.26m-4.28 6.52h4.28m-4.28 6.53h4.28m-15.99 0h6.98m-7.83-10.86L24 27.26l6.67-6.66"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}