package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DrinkToGoTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15.72 2.22a.75.75 0 0 1 1.061 1.06L15.56 4.5h2.19c.318 0 .6.2.707.498l1.25 3.5A.75.75 0 0 1 19 9.5h-1.045l-1.587 10.05c-.21 1.465-1.46 2.45-2.97 2.45h-2.796c-1.511 0-2.761-.985-2.969-2.44L6.045 9.5H5.001a.75.75 0 0 1-.707-1.002l1.25-3.5a.75.75 0 0 1 .707-.498h7.188l2.281-2.28ZM17.222 6H6.78l-.714 2h11.871l-.714-2Z"/>`),
		g.Group(children),
	)
}