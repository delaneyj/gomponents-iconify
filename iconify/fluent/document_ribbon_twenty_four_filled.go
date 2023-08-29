package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentRibbonTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 2v6a2 2 0 0 0 2 2.001h6v10.001a2 2 0 0 1-2 2H9v-3.671a4.5 4.5 0 0 0-5-7.073V4a2 2 0 0 1 2-2h6Zm1.5.5V8a.5.5 0 0 0 .5.5h5.5l-6-6ZM9 15.502a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0Zm-6 3.742v3.044a.71.71 0 0 0 1.212.503L5.5 21.503l1.288 1.288A.71.71 0 0 0 8 22.288v-3.044a4.479 4.479 0 0 1-2.5.758a4.476 4.476 0 0 1-2.5-.758Z"/>`),
		g.Group(children),
	)
}