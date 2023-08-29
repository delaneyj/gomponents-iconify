package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxPlotSolidBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M34 12.34V29a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2h18.57a7.404 7.404 0 0 0 1.063 5H19v14h10V13.427c.103.013.206.024.31.033l1.01.02A7.453 7.453 0 0 0 34 12.34ZM7 26h10V12H7Zm2-7h6v5H9Zm6-2H9v-3h6Zm6-5h4.472c.468.352.98.65 1.528.885V17h-6Zm6 10h-6v-3h6Z" class="clr-i-solid--badged clr-i-solid-path-1--badged"/><circle cx="30" cy="6" r="5" fill="currentColor" class="clr-i-solid--badged clr-i-solid-path-2--badged clr-i-badge"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}