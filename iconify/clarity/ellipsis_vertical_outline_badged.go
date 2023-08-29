package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EllipsisVerticalOutlineBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<circle cx="18" cy="4.9" r="2.9" fill="currentColor" class="clr-i-outline--badged clr-i-outline-path-1--badged"/><circle cx="18" cy="18" r="2.9" fill="currentColor" class="clr-i-outline--badged clr-i-outline-path-2--badged"/><circle cx="18" cy="31.1" r="2.9" fill="currentColor" class="clr-i-outline--badged clr-i-outline-path-3--badged"/><circle cx="30" cy="6" r="5" fill="currentColor" class="clr-i-outline--badged clr-i-outline-path-4--badged clr-i-badge"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}