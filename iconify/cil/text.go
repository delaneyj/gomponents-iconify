package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Text(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M440 168V88H72v80h64v-48h88v280h-56v32h176v-32h-56V120h88v48h64z"/>`),
		g.Group(children),
	)
}