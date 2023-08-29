package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PieChartOutlineBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M32 13.22V29H4V7h18.57a7.447 7.447 0 0 1-.07-1c.001-.335.024-.669.07-1H4a2 2 0 0 0-2 2v22a2 2 0 0 0 2 2h28a2 2 0 0 0 2-2V12.34c-.62.39-1.294.686-2 .88Z" class="clr-i-outline--badged clr-i-outline-path-1--badged"/><path fill="currentColor" d="M17 27a8 8 0 0 1 0-16v8h8a8 8 0 0 1-8 8Zm6.247-6.6H15.4v-7.598A6.4 6.4 0 0 0 17 25.4a6.403 6.403 0 0 0 6.247-5Z" class="clr-i-outline--badged clr-i-outline-path-2--badged"/><path fill="currentColor" d="M19 9a8 8 0 0 1 8 8h-8Zm6.198 6.4a6.409 6.409 0 0 0-4.598-4.599V15.4Z" class="clr-i-outline--badged clr-i-outline-path-3--badged"/><circle cx="30" cy="6" r="5" fill="currentColor" class="clr-i-outline--badged clr-i-outline-path-4--badged clr-i-badge"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}