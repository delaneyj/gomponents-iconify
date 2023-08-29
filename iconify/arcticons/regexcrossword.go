package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Regexcrossword(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15 30.23V17.48h4.18a4.29 4.29 0 0 1 0 8.57H15m4.17 0l4.17 4.18m9.65-8.45l-6.38 8.45m6.38 0l-6.38-8.45M14.6 10.18L7.87 26m32.26-4L33.4 37.82"/>`),
		g.Group(children),
	)
}