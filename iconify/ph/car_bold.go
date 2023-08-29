package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M240 108h-8.2l-26.72-60.12A20 20 0 0 0 186.8 36H69.2a20 20 0 0 0-18.28 11.88L24.2 108H16a12 12 0 0 0 0 24h4v76a20 20 0 0 0 20 20h24a20 20 0 0 0 20-20v-20h88v20a20 20 0 0 0 20 20h24a20 20 0 0 0 20-20v-76h4a12 12 0 0 0 0-24ZM71.8 60h112.4l21.33 48H50.47ZM60 204H44v-16h16Zm136 0v-16h16v16Zm16-40H44v-32h168Z"/>`),
		g.Group(children),
	)
}