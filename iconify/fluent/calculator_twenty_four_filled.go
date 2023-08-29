package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalculatorTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 5.25A3.25 3.25 0 0 1 7.25 2h9.5A3.25 3.25 0 0 1 20 5.25v13.5A3.25 3.25 0 0 1 16.75 22h-9.5A3.25 3.25 0 0 1 4 18.75V5.25ZM9 5a2 2 0 0 0-2 2v1a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2H9Zm.5 8.25a1.25 1.25 0 1 0-2.5 0a1.25 1.25 0 0 0 2.5 0ZM8.25 18.5a1.25 1.25 0 1 0 0-2.5a1.25 1.25 0 0 0 0 2.5ZM17 13.25a1.25 1.25 0 1 0-2.5 0a1.25 1.25 0 0 0 2.5 0Zm-1.25 5.25a1.25 1.25 0 1 0 0-2.5a1.25 1.25 0 0 0 0 2.5Zm-2.5-5.25a1.25 1.25 0 1 0-2.5 0a1.25 1.25 0 0 0 2.5 0ZM12 18.5a1.25 1.25 0 1 0 0-2.5a1.25 1.25 0 0 0 0 2.5Z"/>`),
		g.Group(children),
	)
}