package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Refill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.017 4.479c-1.042 8.421-13.405 16.364-13.405 26.16c0 9.197 7.704 12.84 14.214 12.84s12.483-6.27 12.483-14.453s-8.01-18.515-13.292-24.547Z"/>`),
		g.Group(children),
	)
}