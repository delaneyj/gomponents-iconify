package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dartchecker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.424 45.229a3.5 3.5 0 1 0-6.853-.024"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.886 22.475a23.497 23.497 0 0 0-15.75-.008m14.862 21.868a7 7 0 0 0-13.996 0m4.639-6.422L10.632 7.165m15.727 30.747L37.368 7.165"/>`),
		g.Group(children),
	)
}