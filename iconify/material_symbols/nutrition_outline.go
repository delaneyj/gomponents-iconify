package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NutritionOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21q-2.925 0-4.963-2.038T5 14q0-2.35 1.388-4.213t3.637-2.512q-.5-.125-.975-.363T8.2 6.3q-.825-.825-1.062-1.962t-.113-2.313Q8.2 1.9 9.338 2.138T11.3 3.2q.575.575.838 1.3t.337 1.525q.325-.775.788-1.463T14.3 3.3q.275-.275.7-.275t.7.275q.275.275.275.7t-.275.7q-.55.55-.975 1.213T14.1 7.325q2.2.7 3.55 2.538T19 14q0 2.925-2.038 4.963T12 21Zm0-2q2.075 0 3.538-1.463T17 14q0-2.075-1.463-3.538T12 9q-2.075 0-3.538 1.463T7 14q0 2.075 1.463 3.538T12 19Zm0-5Z"/>`),
		g.Group(children),
	)
}