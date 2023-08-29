package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScatterPlotOutlineBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M32 13.22V29H4V7h18.57a7.447 7.447 0 0 1-.07-1c.001-.335.024-.669.07-1H4a2 2 0 0 0-2 2v22a2 2 0 0 0 2 2h28a2 2 0 0 0 2-2V12.34c-.62.39-1.294.686-2 .88Z" class="clr-i-outline--badged clr-i-outline-path-1--badged"/><path fill="currentColor" d="M9.101 15.8a.8.8 0 0 0 1.13 0l1.16-1.16l1.16 1.16a.8.8 0 1 0 1.13-1.13l-1.15-1.13l1.16-1.16a.8.8 0 1 0-1.13-1.13l-1.16 1.16l-1.17-1.19a.8.8 0 1 0-1.13 1.13l1.16 1.19l-1.16 1.13a.8.8 0 0 0 0 1.13Z" class="clr-i-outline--badged clr-i-outline-path-2--badged"/><path fill="currentColor" d="M15.176 25.536a.8.8 0 0 0 1.13 0l1.16-1.16l1.16 1.16a.8.8 0 1 0 1.13-1.13l-1.15-1.13l1.16-1.16a.8.8 0 1 0-1.13-1.13l-1.16 1.16l-1.17-1.19a.8.8 0 1 0-1.13 1.13l1.16 1.19l-1.16 1.13a.8.8 0 0 0 0 1.13Z" class="clr-i-outline--badged clr-i-outline-path-3--badged"/><path fill="currentColor" d="M22.912 20.343a.8.8 0 0 0 1.13 0l1.16-1.16l1.16 1.16a.8.8 0 1 0 1.13-1.13l-1.15-1.13l1.16-1.16a.8.8 0 1 0-1.13-1.13l-1.16 1.16l-1.17-1.19a.8.8 0 1 0-1.13 1.13l1.16 1.19l-1.16 1.13a.8.8 0 0 0 0 1.13Z" class="clr-i-outline--badged clr-i-outline-path-4--badged"/><circle cx="30" cy="6" r="5" fill="currentColor" class="clr-i-outline--badged clr-i-outline-path-5--badged clr-i-badge"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}