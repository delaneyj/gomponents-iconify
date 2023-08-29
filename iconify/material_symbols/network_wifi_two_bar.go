package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NetworkWifiTwoBar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21L0 9q2.4-2.45 5.5-3.725T12 4q3.425 0 6.525 1.275T24 9L12 21Zm-4.525-7.375q.95-.7 2.1-1.088T12 12.15q1.275 0 2.425.388t2.1 1.087L21.1 9.05q-1.95-1.475-4.263-2.263T12 6q-2.525 0-4.838.788T2.9 9.05l4.575 4.575Z"/>`),
		g.Group(children),
	)
}