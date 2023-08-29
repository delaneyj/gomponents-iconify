package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarSimpleBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M240 108h-8.2l-26.72-60.12A20 20 0 0 0 186.8 36H69.2a20 20 0 0 0-18.28 11.88L24.2 108H16a12 12 0 0 0 0 24h4v76a20 20 0 0 0 20 20h24a20 20 0 0 0 20-20v-12h88v12a20 20 0 0 0 20 20h24a20 20 0 0 0 20-20v-76h4a12 12 0 0 0 0-24ZM71.8 60h112.4l21.33 48H50.47ZM212 204h-16v-20a12 12 0 0 0-12-12H72a12 12 0 0 0-12 12v20H44v-72h168Z"/>`),
		g.Group(children),
	)
}