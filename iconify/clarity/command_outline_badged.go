package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommandOutlineBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M17 23h6v2h-6z" class="clr-i-outline--badged clr-i-outline-path-1--badged"/><path fill="currentColor" d="m7 24.11l9.6-4.41v-1.81L7 13.48v2.2l6.79 3.12L7 21.91v2.2z" class="clr-i-outline--badged clr-i-outline-path-2--badged"/><path fill="currentColor" d="M32 13.22V29H4V10.8h20.24a7.51 7.51 0 0 1-1-1.6H4V7h18.57a7.52 7.52 0 0 1-.07-1a7.52 7.52 0 0 1 .07-1H4a2 2 0 0 0-2 2v22a2 2 0 0 0 2 2h28a2 2 0 0 0 2-2V12.34a7.45 7.45 0 0 1-2 .88Z" class="clr-i-outline--badged clr-i-outline-path-3--badged"/><circle cx="30" cy="6" r="5" fill="currentColor" class="clr-i-outline--badged clr-i-outline-path-4--badged clr-i-badge"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}