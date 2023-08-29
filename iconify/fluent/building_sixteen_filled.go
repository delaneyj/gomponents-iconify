package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.5 2A1.5 1.5 0 0 0 4 3.5v10a.5.5 0 0 0 .5.5H6v-2.5a.5.5 0 0 1 .5-.5h4a.5.5 0 0 1 .5.5V14h1.5a.5.5 0 0 0 .5-.5v-5A1.5 1.5 0 0 0 11.5 7H11V3.5A1.5 1.5 0 0 0 9.5 2h-4ZM7 4.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0ZM7 7a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm0 2.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm2-5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0ZM9 7a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm0 2.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm2 0a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0ZM10 14v-2h-.997v2H10Zm-1.997 0v-2H7v2h1.003Z"/>`),
		g.Group(children),
	)
}