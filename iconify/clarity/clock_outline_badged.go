package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClockOutlineBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M18.92 10.75a1 1 0 0 0-2 0v8.72l5.9 4a1 1 0 1 0 1.11-1.66l-5-3.39Z" class="clr-i-outline--badged clr-i-outline-path-1--badged"/><path fill="currentColor" d="M33.12 12.81a7.44 7.44 0 0 1-1.91.58a14.05 14.05 0 1 1-8.6-8.6a7.44 7.44 0 0 1 .58-1.91a16.06 16.06 0 1 0 9.93 9.93Z" class="clr-i-outline--badged clr-i-outline-path-2--badged"/><path fill="currentColor" d="M18 6.38a11.56 11.56 0 0 0-2.27 22.89L16 27.7a10 10 0 1 1 7.39-18.1a7.45 7.45 0 0 1-.78-2.23A11.45 11.45 0 0 0 18 6.38Z" class="clr-i-outline--badged clr-i-outline-path-3--badged"/><circle cx="30" cy="6" r="5" fill="currentColor" class="clr-i-outline--badged clr-i-outline-path-4--badged clr-i-badge"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}