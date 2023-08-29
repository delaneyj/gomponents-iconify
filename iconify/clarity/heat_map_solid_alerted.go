package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeatMapSolidAlerted(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M34 29a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2h17.958l-2.93 5H8v16h20V15.357h6ZM10 19h4v5h-4Zm12 5v-5h4v5Zm-2-5v5h-4v-5Zm6-2h-4v-1.643h4Zm-6 0h-4v-5h1.856l-.231.395c-.83 1.206-.031 2.85 1.439 2.956c.07.006.137.008.206.006H20Zm-6-5v5h-4v-5Z" class="clr-i-solid--alerted clr-i-solid-path-1--alerted"/><path fill="currentColor" d="m26.854 1.144l-5.72 9.86a1.28 1.28 0 0 0 1.1 2h11.45a1.28 1.28 0 0 0 1.1-2l-5.72-9.86a1.28 1.28 0 0 0-2.21 0Z" class="clr-i-solid--alerted clr-i-solid-path-22--alerted clr-i-alert"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}