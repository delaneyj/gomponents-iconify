package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DisplaySolidBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M26 32h-1.74a3.61 3.61 0 0 1-1.5-2.52v-1.35h-9.52v1.37a3.61 3.61 0 0 1-1.5 2.5h-1.8a1 1 0 1 0 0 2h16.12a.92.92 0 0 0 1-1A1 1 0 0 0 26 32Z" class="clr-i-solid--badged clr-i-solid-path-1--badged"/><path fill="currentColor" d="M31 13.43v8.41H5V7h17.57a7.29 7.29 0 0 1 .55-4H3.5A1.5 1.5 0 0 0 2 4.5v21A1.5 1.5 0 0 0 3.5 27h29a1.5 1.5 0 0 0 1.5-1.5V12.34a7.44 7.44 0 0 1-3 1.09Z" class="clr-i-solid--badged clr-i-solid-path-2--badged"/><circle cx="30" cy="6" r="5" fill="currentColor" class="clr-i-solid--badged clr-i-solid-path-3--badged clr-i-badge"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}