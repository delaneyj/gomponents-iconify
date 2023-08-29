package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeMoreTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.44 2.532a2.25 2.25 0 0 0-2.903.002l-6.74 5.698A2.25 2.25 0 0 0 3 9.95v9.3c0 .966.784 1.75 1.75 1.75h14.5A1.75 1.75 0 0 0 21 19.25v-9.3a2.25 2.25 0 0 0-.8-1.72l-6.76-5.7ZM9 11.75a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0Zm4.25 0a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0Zm4.25 0a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0Zm-8.5 4a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0ZM12 17a1.25 1.25 0 1 1 0-2.5a1.25 1.25 0 0 1 0 2.5Zm4.25 0a1.25 1.25 0 1 1 0-2.5a1.25 1.25 0 0 1 0 2.5Z"/>`),
		g.Group(children),
	)
}