package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserSolidBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M30.61 24.52a17.16 17.16 0 0 0-25.22 0a1.51 1.51 0 0 0-.39 1v6A1.5 1.5 0 0 0 6.5 33h23a1.5 1.5 0 0 0 1.5-1.5v-6a1.51 1.51 0 0 0-.39-.98Z" class="clr-i-solid--badged clr-i-solid-path-1--badged"/><path fill="currentColor" d="M18 17a7 7 0 0 0 6.85-5.56a7.4 7.4 0 0 1-2.24-6.69A7 7 0 1 0 18 17Z" class="clr-i-solid--badged clr-i-solid-path-2--badged"/><circle cx="30" cy="6" r="5" fill="currentColor" class="clr-i-solid--badged clr-i-solid-path-3--badged clr-i-badge"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}