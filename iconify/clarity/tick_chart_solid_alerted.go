package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TickChartSolidAlerted(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M34 29a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2h17.958l-4.333 7.395c-.83 1.206-.031 2.85 1.439 2.956c.07.006.137.008.206.006H34ZM6 26.991l21.723-.001c.76-.06 1.18-.93.74-1.56a.974.974 0 0 0-.74-.43H25v-3h-2v3h-3v-3h-2v3h-3v-3h-2v3h-3v-3H8v3H6Z" class="clr-i-solid--alerted clr-i-solid-path-1--alerted"/><path fill="currentColor" d="m26.854 1.144l-5.72 9.86a1.28 1.28 0 0 0 1.1 2h11.45a1.28 1.28 0 0 0 1.1-2l-5.72-9.86a1.28 1.28 0 0 0-2.21 0Z" class="clr-i-solid--alerted clr-i-solid-path-2--alerted clr-i-alert"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}