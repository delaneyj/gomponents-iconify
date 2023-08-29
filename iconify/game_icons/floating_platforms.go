package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FloatingPlatforms(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M184 43.5v18h144l-.1-18zm24 95.9l-144 .1v18l144-.1zm96.1 0v18l143.9.1v-18.1zm168.4 98.1l-435.63.1L128 372.5l10.9-32l21.1 80l32-52.5l48 128l80-107.5l32 16z"/>`),
		g.Group(children),
	)
}