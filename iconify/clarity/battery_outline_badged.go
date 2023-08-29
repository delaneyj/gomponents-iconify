package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryOutlineBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<circle cx="30" cy="6" r="5" fill="currentColor" class="clr-i-outline--badged clr-i-outline-path-1--badged clr-i-badge"/><path fill="currentColor" d="M18.59 11.77a1 1 0 0 0-1.73 1l2.5 4.34l-6.07-1l5.29 10.59a1 1 0 0 0 1.79-.89l-3.53-7.08l6.38 1.06Z" class="clr-i-outline--badged clr-i-outline-path-2--badged"/><path fill="currentColor" d="M25 11.58V32H11V6h4V4h6v2h1.5a7.47 7.47 0 0 1 .5-2.62A1.57 1.57 0 0 0 21.42 2h-6.84A1.58 1.58 0 0 0 13 3.58V4h-2.12A1.88 1.88 0 0 0 9 5.88v26.24A1.88 1.88 0 0 0 10.88 34h14.24A1.88 1.88 0 0 0 27 32.12V12.87a7.5 7.5 0 0 1-2-1.29Z" class="clr-i-outline--badged clr-i-outline-path-3--badged"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}