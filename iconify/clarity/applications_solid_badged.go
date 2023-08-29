package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ApplicationsSolidBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M4 4h6v6H4z" class="clr-i-solid--badged clr-i-solid-path-1--badged"/><path fill="currentColor" d="M4 15h6v6H4z" class="clr-i-solid--badged clr-i-solid-path-2--badged"/><path fill="currentColor" d="M4 26h6v6H4z" class="clr-i-solid--badged clr-i-solid-path-3--badged"/><path fill="currentColor" d="M15 4h6v6h-6z" class="clr-i-solid--badged clr-i-solid-path-4--badged"/><path fill="currentColor" d="M15 15h6v6h-6z" class="clr-i-solid--badged clr-i-solid-path-5--badged"/><path fill="currentColor" d="M15 26h6v6h-6z" class="clr-i-solid--badged clr-i-solid-path-6--badged"/><path fill="currentColor" d="M26 15h6v6h-6z" class="clr-i-solid--badged clr-i-solid-path-7--badged"/><path fill="currentColor" d="M26 26h6v6h-6z" class="clr-i-solid--badged clr-i-solid-path-8--badged"/><circle cx="30" cy="6" r="5" fill="currentColor" class="clr-i-solid--badged clr-i-solid-path-9--badged clr-i-badge"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}