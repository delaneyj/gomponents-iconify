package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Swimming(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M10.111 2c-.112 0-.435.147-.435.147l-3.322 1.68c-.443.175-.618.881-.352 1.234l.97 1.408l-3.97 2.029L5 9.998l2.502-1.5l2.5 1.5l1.002-1.002l-3-4l2.557-1.53c.528-.266.443-.704.443-.97C11 2.286 10.644 2 10.111 2zm2.141 3a1.75 1.75 0 1 0-.003 3.501A1.75 1.75 0 0 0 12.252 5zM2.5 10L0 11.5V13l2.5-1.5L5 13l2.502-1.5l2.5 1.5L12 11.5l3 1.5v-1.5L12 10l-1.998 1.5l-2.5-1.5L5 11.5L2.5 10z"/>`),
		g.Group(children),
	)
}