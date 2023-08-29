package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PowerOutlineBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M18 21a1 1 0 0 1-1-1V4a1 1 0 0 1 2 0v16a1 1 0 0 1-1 1Z" class="clr-i-outline--badged clr-i-outline-path-1--badged"/><path fill="currentColor" d="M30 13.5h-.31a13 13 0 1 1-18.21-5.6a1 1 0 0 0-1-1.73a15 15 0 1 0 21.31 7.1a7.49 7.49 0 0 1-1.79.23Z" class="clr-i-outline--badged clr-i-outline-path-2--badged"/><circle cx="30" cy="6" r="5" fill="currentColor" class="clr-i-outline--badged clr-i-outline-path-3--badged clr-i-badge"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}