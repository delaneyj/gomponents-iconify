package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineHorizontalFourSearchSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.5 8c.786 0 1.512-.26 2.096-.697l2.55 2.55a.5.5 0 1 0 .708-.707l-2.55-2.55A3.5 3.5 0 1 0 10.5 8Zm0-1a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5Zm3.288 2.909L12.878 9H2.5a.5.5 0 0 0 0 1h11a.498.498 0 0 0 .288-.091ZM6.758 7a4.484 4.484 0 0 1-.502-1H2.5a.5.5 0 0 0 0 1h4.258Zm-.502-4c-.112.318-.19.653-.229 1H2.5a.5.5 0 0 1 0-1h3.756ZM2.5 12a.5.5 0 0 0 0 1h11a.5.5 0 0 0 0-1h-11Z"/>`),
		g.Group(children),
	)
}