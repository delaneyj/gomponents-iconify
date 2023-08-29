package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JeepBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M240 92h-6.3l-8.61-40.19A20.11 20.11 0 0 0 205.53 36H50.47a20.11 20.11 0 0 0-19.56 15.81L22.3 92H16a12 12 0 0 0 0 24h4v92a20 20 0 0 0 20 20h24a20 20 0 0 0 20-20v-20h88v20a20 20 0 0 0 20 20h24a20 20 0 0 0 20-20v-92h4a12 12 0 0 0 0-24ZM53.7 60h148.6l6.86 32H46.84ZM60 204H44v-16h16Zm136 0v-16h16v16Zm16-40h-32v-24a12 12 0 0 0-24 0v24h-16v-24a12 12 0 0 0-24 0v24h-16v-24a12 12 0 0 0-24 0v24H44v-48h168Z"/>`),
		g.Group(children),
	)
}