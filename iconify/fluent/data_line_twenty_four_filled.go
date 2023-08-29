package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataLineTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16 6a3 3 0 1 1 2.524 2.962l-2.038 3.358a3 3 0 0 1-4.749 3.65l-3.741 1.87a3 3 0 1 1-.461-1.446l3.531-1.765a3 3 0 0 1 4.275-3.313l1.798-2.963A2.995 2.995 0 0 1 16 6Z"/>`),
		g.Group(children),
	)
}