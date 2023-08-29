package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bottle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M576 960q-5 26-21.5 45t-42.5 19H128q-26 0-42-18.5T64 960L0 576V448q0-29 19.5-48T77 362.5t47-24.5q40-28 72-66t47-80h-19q-13 0-22.5-9.5T192 160V32q0-13 9.5-22.5T224 0h192q13 0 22.5 9.5T448 32v128q0 13-9.5 22.5T416 192h-19q30 81 111 140q11 9 50 28t60.5 39t21.5 49v128z"/>`),
		g.Group(children),
	)
}