package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldSolidAlerted(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M22.23 15.4A3.68 3.68 0 0 1 19 9.89L22.43 4a41.1 41.1 0 0 1-3.89-2.18L18 1.5l-.54.35a45 45 0 0 1-6.08 3.21A43.79 43.79 0 0 1 4.75 7.4L4 7.59v8.34c0 13.39 13.53 18.4 13.66 18.45l.34.12l.34-.12c.14 0 13.66-5.05 13.66-18.45v-.53Z" class="clr-i-solid--alerted clr-i-solid-path-1--alerted"/><path fill="currentColor" d="M26.85 1.14L21.13 11a1.28 1.28 0 0 0 1.1 2h11.45a1.28 1.28 0 0 0 1.1-2l-5.72-9.86a1.28 1.28 0 0 0-2.21 0Z" class="clr-i-solid--alerted clr-i-solid-path-2--alerted clr-i-alert"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}