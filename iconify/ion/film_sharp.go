package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilmSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M480 80H32v352h448ZM112 352v48H64v-48Zm0-80v48H64v-48Zm0-80v48H64v-48Zm0-80v48H64v-48Zm256 160H144v-32h224Zm80 80v48h-48v-48Zm0-80v48h-48v-48Zm0-80v48h-48v-48Zm0-80v48h-48v-48Z"/>`),
		g.Group(children),
	)
}