package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HourglassSolidBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M28.67 32h-22a1 1 0 0 0 0 2h22a1 1 0 1 0 0-2Z" class="clr-i-solid--badged clr-i-solid-path-1--badged"/><path fill="currentColor" d="M6.67 4h16.11a7.45 7.45 0 0 1 .89-2h-17a1 1 0 1 0 0 2Z" class="clr-i-solid--badged clr-i-solid-path-2--badged"/><path fill="currentColor" d="M22.55 20.27a11.48 11.48 0 0 0-2.91-1.72v-1.16a11.48 11.48 0 0 0 2.91-1.72A6.25 6.25 0 0 0 25 11.55A7.47 7.47 0 0 1 22.5 6H10.06v5.12a6.07 6.07 0 0 0 2.45 4.55a11.48 11.48 0 0 0 2.91 1.72v1.16a11.48 11.48 0 0 0-2.91 1.72a6.07 6.07 0 0 0-2.45 4.55v5.12H25v-5.12a6.07 6.07 0 0 0-2.45-4.55Z" class="clr-i-solid--badged clr-i-solid-path-3--badged"/><circle cx="30" cy="6" r="5" fill="currentColor" class="clr-i-solid--badged clr-i-solid-path-4--badged clr-i-badge"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}