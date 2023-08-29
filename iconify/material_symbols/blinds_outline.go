package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlindsOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20v-2h2V4h16v14h2v2ZM6 9h12V6H6Zm0 9h12v-7h-5v2.275q.45.275.725.725q.275.45.275 1q0 .825-.587 1.413Q12.825 17 12 17q-.825 0-1.412-.587Q10 15.825 10 15q0-.55.275-1q.275-.45.725-.725V11H6ZM6 6h12H6Z"/>`),
		g.Group(children),
	)
}