package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JBrick(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M265.5 83.5A9.5 9.5 0 0 0 256 93v96a9.5 9.5 0 0 0 9.5 9.5a9.5 9.5 0 0 0-9.5 9.5v96a9.5 9.5 0 0 0 9.5 9.5a9.5 9.5 0 0 0-9.5 9.5a9.5 9.5 0 0 0-9.5-9.5h-96a9.5 9.5 0 0 0-9.5 9.5v96a9.5 9.5 0 0 0 9.5 9.5h96a9.5 9.5 0 0 0 9.5-9.5a9.5 9.5 0 0 0 9.5 9.5h96a9.5 9.5 0 0 0 9.5-9.5v-96a9.5 9.5 0 0 0-9.5-9.5a9.5 9.5 0 0 0 9.5-9.5v-96a9.5 9.5 0 0 0-9.5-9.5a9.5 9.5 0 0 0 9.5-9.5V93a9.5 9.5 0 0 0-9.5-9.5h-96zm9.5 19h77v77h-77v-77zm0 115h77v77h-77v-77zm-115 115h77v77h-77v-77zm115 0h77v77h-77v-77z"/>`),
		g.Group(children),
	)
}