package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClothesBriefs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M22.158 37s-.9-8.075-4.158-12c-3.044-3.669-12-6-12-6v-5h36v5s-8.956 2.331-12 6c-3.257 3.925-4.158 12-4.158 12h-3.684Z"/>`),
		g.Group(children),
	)
}