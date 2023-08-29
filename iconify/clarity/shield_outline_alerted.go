package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldOutlineAlerted(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M30 15.4v.53c0 11-10 15.61-12 16.43c-2-.82-12-5.44-12-16.43V9.14a47.54 47.54 0 0 0 6.18-2.25a48.23 48.23 0 0 0 5.82-3c1 .64 2.2 1.27 3.43 1.89l1-1.74a41.1 41.1 0 0 1-3.89-2.18L18 1.5l-.54.35a45 45 0 0 1-6.08 3.21A43.79 43.79 0 0 1 4.75 7.4L4 7.59v8.34c0 13.39 13.53 18.4 13.66 18.45l.34.12l.34-.12c.14 0 13.66-5.05 13.66-18.45v-.53Z" class="clr-i-outline--alerted clr-i-outline-path-1--alerted"/><path fill="currentColor" d="M26.85 1.14L21.13 11a1.28 1.28 0 0 0 1.1 2h11.45a1.28 1.28 0 0 0 1.1-2l-5.72-9.86a1.28 1.28 0 0 0-2.21 0Z" class="clr-i-outline--alerted clr-i-outline-path-2--alerted clr-i-alert"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}