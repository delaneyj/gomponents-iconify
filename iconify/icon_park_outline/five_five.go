package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FiveFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M14.223 37.471L6.55 26.984a4.423 4.423 0 0 1 6.696-5.738L16 24V7.25a3.25 3.25 0 0 1 6.5 0v-1a3.25 3.25 0 0 1 6.5 0v1a3.25 3.25 0 0 1 6.5 0v4a3.25 3.25 0 0 1 6.5 0v19.058c0 2.73-.838 5.417-2.38 7.671C37.056 41.726 32.785 44 28.245 44H27.09a15.944 15.944 0 0 1-12.867-6.529Z"/>`),
		g.Group(children),
	)
}