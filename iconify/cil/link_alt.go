package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M48 256a80.091 80.091 0 0 1 80-80h112v-32H128a112 112 0 0 0 0 224h112v-32H128a80.091 80.091 0 0 1-80-80Zm336-112H272v32h112a80 80 0 0 1 0 160H272v32h112a112 112 0 0 0 0-224Z"/><path fill="currentColor" d="M140 240.652h232v32H140z"/>`),
		g.Group(children),
	)
}