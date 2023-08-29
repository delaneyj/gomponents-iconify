package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagClockFortyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.5 7.25C7.5 6.56 8.06 6 8.75 6h32.5a1.25 1.25 0 0 1 1.007 1.99L33.801 19.5l1.849 2.516a13.075 13.075 0 0 0-2.963.19l-1.444-1.966a1.25 1.25 0 0 1 0-1.48L38.78 8.5H10v22h12.8a12.89 12.89 0 0 0-.647 2.5H10v9.75a1.25 1.25 0 1 1-2.5 0V7.25ZM46 35c0 6.075-4.925 11-11 11s-11-4.925-11-11s4.925-11 11-11s11 4.925 11 11Zm-5 0a1 1 0 0 0-1-1h-4v-6a1 1 0 1 0-2 0v7a1 1 0 0 0 1 1h5a1 1 0 0 0 1-1Z"/>`),
		g.Group(children),
	)
}