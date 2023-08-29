package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Windows(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m3 5.548l7.195-.966v7.029l-7.188.054L3 5.55Zm7.195 6.843v7.105l-7.19-.985v-6.12h7.19Zm.918-7.935L20.998 3v8.533l-9.885.078V4.456ZM21 12.505L20.998 21l-9.885-1.353v-7.142H21Z"/>`),
		g.Group(children),
	)
}