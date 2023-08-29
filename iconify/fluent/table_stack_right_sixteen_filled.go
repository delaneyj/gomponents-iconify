package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableStackRightSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14 2.5a.5.5 0 0 0-1 0v11a.5.5 0 0 0 1 0v-11ZM10 11v2.5a.5.5 0 0 1-.5.5H6v-3h4Zm0-1V6H6v4h4ZM6 2h3.5a.5.5 0 0 1 .5.5V5H6V2ZM5 5V2h-.5A2.5 2.5 0 0 0 2 4.5V5h3ZM2 6v4h3V6H2Zm0 5.5V11h3v3h-.5A2.5 2.5 0 0 1 2 11.5Z"/>`),
		g.Group(children),
	)
}