package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MemoryOutlineBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M8 12h4v8H8z" class="clr-i-outline--badged clr-i-outline-path-1--badged"/><path fill="currentColor" d="M16 12h4v8h-4z" class="clr-i-outline--badged clr-i-outline-path-2--badged"/><path fill="currentColor" d="M15 27H4V17H2v10a2 2 0 0 0 2 2h12.61v-3.45h2.26V24H15Z" class="clr-i-outline--badged clr-i-outline-path-3--badged"/><path fill="currentColor" d="M32 17v10H19v2h13a2 2 0 0 0 2-2V17Z" class="clr-i-outline--badged clr-i-outline-path-4--badged"/><path fill="currentColor" d="M28 13.22A7.46 7.46 0 0 1 25.51 12H24v8h4Z" class="clr-i-outline--badged clr-i-outline-path-5--badged"/><path fill="currentColor" d="M4 9h19.13a7.45 7.45 0 0 1-.55-2H4a2 2 0 0 0-2 2v4h2Z" class="clr-i-outline--badged clr-i-outline-path-6--badged"/><circle cx="30" cy="6" r="5" fill="currentColor" class="clr-i-outline--badged clr-i-outline-path-7--badged clr-i-badge"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}