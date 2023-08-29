package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VmOutlineBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M11 5h11.57a7.45 7.45 0 0 1 .55-2H11a2 2 0 0 0-2 2v6.85h2Z" class="clr-i-outline--badged clr-i-outline-path-1--badged"/><path fill="currentColor" d="M30 13.5V26h-8v-9a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-3h8a2 2 0 0 0 2-2V13.22a7.49 7.49 0 0 1-2 .28ZM6 31V17h14v9h-4v-6h-2v6a2 2 0 0 0 2 2h4v3Z" class="clr-i-outline--badged clr-i-outline-path-2--badged"/><path fill="currentColor" d="M17 12h8v6h2v-5.13A7.52 7.52 0 0 1 23.66 10H17Z" class="clr-i-outline--badged clr-i-outline-path-3--badged"/><circle cx="30" cy="6" r="5" fill="currentColor" class="clr-i-outline--badged clr-i-outline-path-4--badged clr-i-badge"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}