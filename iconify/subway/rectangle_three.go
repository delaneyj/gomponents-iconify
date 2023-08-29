package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RectangleThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M139.6 46.5h93.1V0h-93.1v46.5zM418.9 0v46.5h46.5V93H512V0h-93.1zM46.5 46.5H93V0H0v93.1h46.5V46.5zM279.3 512h93.1v-46.5h-93.1V512zM46.5 139.6H0v93.1h46.5v-93.1zm419 93.1H512v-93.1h-46.5v93.1zm0 139.7H512v-93.1h-46.5v93.1zm0 93.1H419V512h93v-93.1h-46.5v46.6zm-186.2-419h93.1V0h-93.1v46.5zM0 512h232.7V279.3H0V512z"/>`),
		g.Group(children),
	)
}