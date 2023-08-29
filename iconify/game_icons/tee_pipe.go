package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TeePipe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M25 115v154h30V115zm432 0v154h30V115zM73 128v128h103.8l40-53.4l14.4 10.8l-39.2 52.3V439h128V265.7l-39.2-52.3l14.4-10.8l40 53.4H439V128zm23 23h320v18H96zm119 137h18v128h-18zm-36 169v30h154v-30z"/>`),
		g.Group(children),
	)
}