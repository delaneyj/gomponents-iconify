package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExportSolidBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M29.32 16.35a1 1 0 0 0-1.41 1.41L31.16 21H26v2h5.19l-3.28 3.28a1 1 0 1 0 1.41 1.41L35 22Z" class="clr-i-solid--badged clr-i-solid-path-1--badged"/><path fill="currentColor" d="M17 22a1 1 0 0 1 1-1h8v-8.66A7.46 7.46 0 0 1 22.78 4H10.87L4 10.86V30a2 2 0 0 0 2 2h18a2 2 0 0 0 2-2v-7h-8a1 1 0 0 1-1-1Zm-5-10H6v-.32L11.69 6H12Z" class="clr-i-solid--badged clr-i-solid-path-2--badged"/><circle cx="30" cy="6" r="5" fill="currentColor" class="clr-i-solid--badged clr-i-solid-path-3--badged clr-i-badge"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}