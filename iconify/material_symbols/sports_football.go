package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SportsFootball(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.075 20.975q-1.175.125-2.837-.063T3.7 20.3q-.35-.8-.588-2.5T3 14.9l6.075 6.075Zm2.375-.4l-8.05-8.05q.425-1.875 1.238-3.413t1.887-2.637Q7.6 5.4 9.138 4.613t3.312-1.188l8.1 8.1q-.4 1.85-1.175 3.4t-1.85 2.625q-1.125 1.1-2.688 1.888t-3.387 1.137ZM9.4 16L16 9.4L14.6 8L8 14.6L9.4 16Zm11.55-6.85l-6.075-6.125q1.2-.15 2.95.05T20.3 3.7q.45 1 .625 2.588t.025 2.862Z"/>`),
		g.Group(children),
	)
}