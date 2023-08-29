package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TickChartOutlineBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M32 13.22V29H4V7h18.57a7.447 7.447 0 0 1-.07-1c.001-.335.024-.669.07-1H4a2 2 0 0 0-2 2v22a2 2 0 0 0 2 2h28a2 2 0 0 0 2-2V12.34c-.62.39-1.294.686-2 .88Z" class="clr-i-outline--badged clr-i-outline-path-1--badged"/><path fill="currentColor" d="M6 25h2v-3h2v3h3v-3h2v3h3v-3h2v3h3v-3h2v3h2.723c.3.02.57.18.74.43c.44.63.02 1.5-.74 1.56L6 26.991Z" class="clr-i-outline--badged clr-i-outline-path-2--badged"/><circle cx="30" cy="6" r="5" fill="currentColor" class="clr-i-outline--badged clr-i-outline-path-3--badged clr-i-badge"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}