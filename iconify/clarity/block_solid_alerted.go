package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlockSolidAlerted(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M30 15.38v11l-11 5.06v-14.8l2.79-1.29a3.68 3.68 0 0 1-2.25-1.16L18 14.9L7.39 10L18 5.1l3 1.39l1-1.75l-3.58-1.65a1 1 0 0 0-.84 0l-13 6A1 1 0 0 0 4 10v17a1 1 0 0 0 .58.91l13 6a1 1 0 0 0 .84 0l13-6A1 1 0 0 0 32 27V15.38Z" class="clr-i-solid--alerted clr-i-solid-path-1--alerted"/><path fill="currentColor" d="M26.85 1.12L21.13 11a1.27 1.27 0 0 0 1.1 1.91h11.45a1.27 1.27 0 0 0 1.1-1.91l-5.72-9.88a1.28 1.28 0 0 0-2.21 0Z" class="clr-i-solid--alerted clr-i-solid-path-2--alerted clr-i-alert"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}