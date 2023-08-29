package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarChartOutlineBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M32 13.22V29H4V7h18.57a7.447 7.447 0 0 1-.07-1c.001-.335.024-.669.07-1H4a2 2 0 0 0-2 2v22a2 2 0 0 0 2 2h28a2 2 0 0 0 2-2V12.34c-.62.39-1.294.686-2 .88Z" class="clr-i-outline--badged clr-i-outline-path-1--badged"/><path fill="currentColor" d="M7 10h6v16h-1.6V11.6H8.6V26H7Z" class="clr-i-outline--badged clr-i-outline-path-2--badged"/><path fill="currentColor" d="M15 19h6v7h-1.6v-5.4h-2.8V26H15Z" class="clr-i-outline--badged clr-i-outline-path-3--badged"/><path fill="currentColor" d="M23 16h6v10h-1.6v-8.4h-2.8V26H23Z" class="clr-i-outline--badged clr-i-outline-path-4--badged"/><circle cx="30" cy="6" r="5" fill="currentColor" class="clr-i-outline--badged clr-i-outline-path-5--badged clr-i-badge"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}