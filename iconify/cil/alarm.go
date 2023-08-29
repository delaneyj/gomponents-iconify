package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alarm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 80C141.125 80 48 173.125 48 288s93.125 208 208 208s208-93.125 208-208S370.875 80 256 80Zm124.451 332.451A176 176 0 1 1 432 288a174.849 174.849 0 0 1-51.549 124.451Z"/><path fill="currentColor" d="M272 160h-32v135.69l86.005 68.804l19.99-24.988L272 280.31V160zM22.965 114.796l152-104l18.071 26.411l-152 104zm296.002-77.59l18.07-26.41l152 104.002l-18.071 26.41z"/>`),
		g.Group(children),
	)
}