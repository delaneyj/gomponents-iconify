package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlockOutlineAlerted(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M30 15.53v10.83l-11 5.08v-14.8l2.57-1.19a3.67 3.67 0 0 1-2.11-1.22L18 14.9L7.39 10L18 5.1l3.08 1.42l1-1.74l-3.66-1.69a1 1 0 0 0-.84 0l-13 6A1 1 0 0 0 4 10v17a1 1 0 0 0 .58.91l13 6a1 1 0 0 0 .84 0l13-6A1 1 0 0 0 32 27V15.53ZM17 31.44L6 26.36v-14.8l11 5.08Z" class="clr-i-outline--alerted clr-i-outline-path-1--alerted"/><path fill="currentColor" d="m26.87 1.26l-5.72 9.91a1.28 1.28 0 0 0 1.1 1.92H33.7a1.28 1.28 0 0 0 1.1-1.92l-5.72-9.91a1.28 1.28 0 0 0-2.21 0Z" class="clr-i-outline--alerted clr-i-outline-path-2--alerted clr-i-alert"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}