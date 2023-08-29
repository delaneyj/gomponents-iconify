package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RectangleTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M46.5 279.3H0v93.1h46.5v-93.1zm0 139.6H0V512h93.1v-46.5H46.5v-46.6zm93.1 93.1h93.1v-46.5h-93.1V512zM46.5 139.6H0v93.1h46.5v-93.1zM0 93.1h46.5V46.5H93V0H0v93.1zm465.5 372.4H419V512h93v-93.1h-46.5v46.6zm0-93.1H512v-93.1h-46.5v93.1zM139.6 46.5h93.1V0h-93.1v46.5zM279.3 512h93.1v-46.5h-93.1V512zm0-512v232.7H512V0H279.3z"/>`),
		g.Group(children),
	)
}