package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScaleFitSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 4.998v6a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2v-6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2Zm7.65 4.86a.5.5 0 0 1 0-.707l.65-.65H8.752a.5.5 0 0 1 0-1h1.542l-.643-.643a.5.5 0 1 1 .707-.707l1.5 1.5a.5.5 0 0 1 0 .707l-1.5 1.5a.5.5 0 0 1-.707 0ZM6.352 6.151a.5.5 0 0 1 0 .707L5.71 7.5h1.54a.5.5 0 1 1 0 1H5.7l.652.651a.5.5 0 1 1-.707.707l-1.5-1.5a.5.5 0 0 1 0-.707l1.5-1.5a.5.5 0 0 1 .707 0Z"/>`),
		g.Group(children),
	)
}