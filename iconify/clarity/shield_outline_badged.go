package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldOutlineBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M30 13.5v2.43c0 11-10 15.61-12 16.43c-2-.82-12-5.44-12-16.43V9.14a47.54 47.54 0 0 0 6.18-2.25a48.23 48.23 0 0 0 5.82-3a46.19 46.19 0 0 0 4.51 2.42v-.29a7.49 7.49 0 0 1 .23-1.83a41.61 41.61 0 0 1-4.19-2.33L18 1.5l-.54.35a45 45 0 0 1-6.08 3.21A43.79 43.79 0 0 1 4.75 7.4L4 7.59v8.34c0 13.39 13.53 18.4 13.66 18.45l.34.12l.34-.12c.14 0 13.66-5.05 13.66-18.45v-2.71a7.49 7.49 0 0 1-2 .28Z" class="clr-i-outline--badged clr-i-outline-path-1--badged"/><circle cx="30" cy="6" r="5" fill="currentColor" class="clr-i-outline--badged clr-i-outline-path-2--badged clr-i-badge"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}