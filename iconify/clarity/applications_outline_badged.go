package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ApplicationsOutlineBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M8 8H4v2h6V4H8v4z" class="clr-i-outline--badged clr-i-outline-path-1--badged"/><path fill="currentColor" d="M19 8h-4v2h6V4h-2v4z" class="clr-i-outline--badged clr-i-outline-path-2--badged"/><path fill="currentColor" d="M8 19H4v2h6v-6H8v4z" class="clr-i-outline--badged clr-i-outline-path-3--badged"/><path fill="currentColor" d="M19 19h-4v2h6v-6h-2v4z" class="clr-i-outline--badged clr-i-outline-path-4--badged"/><path fill="currentColor" d="M30 19h-4v2h6v-6h-2v4z" class="clr-i-outline--badged clr-i-outline-path-5--badged"/><path fill="currentColor" d="M8 30H4v2h6v-6H8v4z" class="clr-i-outline--badged clr-i-outline-path-6--badged"/><path fill="currentColor" d="M19 30h-4v2h6v-6h-2v4z" class="clr-i-outline--badged clr-i-outline-path-7--badged"/><path fill="currentColor" d="M30 30h-4v2h6v-6h-2v4z" class="clr-i-outline--badged clr-i-outline-path-8--badged"/><circle cx="30" cy="6" r="5" fill="currentColor" class="clr-i-outline--badged clr-i-outline-path-9--badged clr-i-badge"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}