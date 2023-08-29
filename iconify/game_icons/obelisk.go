package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Obelisk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 22.127L238.562 57h34.875L256 22.127zM233 75v268h46V75h-46zm-16 286v14h78v-14h-78zm-17.193 32l-13.43 94h139.246l-13.428-94H199.807z"/>`),
		g.Group(children),
	)
}