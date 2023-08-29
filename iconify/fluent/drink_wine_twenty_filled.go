package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DrinkWineTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.5 2a.5.5 0 0 0-.5.5V8a5 5 0 0 0 4.5 4.975V17h-2a.5.5 0 0 0 0 1h5a.5.5 0 0 0 0-1h-2v-4.025A5 5 0 0 0 15 8V2.5a.5.5 0 0 0-.5-.5h-9ZM6 5.5V3h8v2.5H6Zm7.452 2.575a3.5 3.5 0 0 1-2.867 2.876a.5.5 0 1 1-.167-.986a2.5 2.5 0 0 0 2.048-2.054a.5.5 0 0 1 .986.164Z"/>`),
		g.Group(children),
	)
}