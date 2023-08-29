package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PieChartOutlineAlerted(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M34 29a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2h17.958l-1.172 2H4v22h28V15.357h2Z" class="clr-i-outline--alerted clr-i-outline-path-1--alerted"/><path fill="currentColor" d="M17 27a8 8 0 0 1 0-16v8h8a8 8 0 0 1-8 8Zm6.247-6.6H15.4v-7.598A6.4 6.4 0 0 0 17 25.4a6.403 6.403 0 0 0 6.247-5Z" class="clr-i-outline--alerted clr-i-outline-path-2--alerted"/><path fill="currentColor" d="M27 17h-8v-1.655c.021.003.043.004.064.006c.07.006.137.008.206.006h7.561c.111.53.169 1.08.169 1.643Z" class="clr-i-outline--alerted clr-i-outline-path-3--alerted"/><path fill="currentColor" d="m26.854 1.144l-5.72 9.86a1.28 1.28 0 0 0 1.1 2h11.45a1.28 1.28 0 0 0 1.1-2l-5.72-9.86a1.28 1.28 0 0 0-2.21 0Z" class="clr-i-outline--alerted clr-i-outline-path-4--alerted clr-i-alert"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}