package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stairs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M64 448v-64h64v-64h64v-64h64v-64h64v-64h64V64h64v384z"/>`),
		g.Group(children),
	)
}