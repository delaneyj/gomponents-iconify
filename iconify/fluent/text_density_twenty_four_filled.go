package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextDensityTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12.25 2a.75.75 0 0 1 .75.75v18.5a.75.75 0 0 1-1.5 0V2.75a.75.75 0 0 1 .75-.75ZM10 5H2.75a.75.75 0 0 0 0 1.5H10V5Zm0 4H2.75a.75.75 0 0 0 0 1.5H10V9Zm0 4H2.75a.75.75 0 0 0 0 1.5H10V13Zm0 4H2.75a.75.75 0 0 0 0 1.5H10V17Zm9.75 1H14.5v-4.5h5.25a2.25 2.25 0 0 1 0 4.5Zm0-7.5H14.5V6h5.25a2.25 2.25 0 0 1 0 4.5Z"/>`),
		g.Group(children),
	)
}