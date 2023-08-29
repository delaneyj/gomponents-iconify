package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HardDiskOutlineBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M6 20h24v2H6z" class="clr-i-outline--badged clr-i-outline-path-1--badged"/><path fill="currentColor" d="M26 24h4v2h-4z" class="clr-i-outline--badged clr-i-outline-path-2--badged"/><path fill="currentColor" d="m34 21.08l-2-7.87a7.49 7.49 0 0 1-2 .29l2 7.94V29H4v-7.56L7.06 9h16.07a7.45 7.45 0 0 1-.55-2H7.06a2 2 0 0 0-1.93 1.47L2 21.08a1 1 0 0 0 0 .24V29a2 2 0 0 0 2 2h28a2 2 0 0 0 2-2v-7.69a1 1 0 0 0 0-.23Z" class="clr-i-outline--badged clr-i-outline-path-3--badged"/><circle cx="30" cy="6" r="5" fill="currentColor" class="clr-i-outline--badged clr-i-outline-path-4--badged clr-i-badge"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}