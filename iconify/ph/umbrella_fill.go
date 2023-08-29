package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UmbrellaFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M240 126.63A112.21 112.21 0 0 0 128 24A112.21 112.21 0 0 0 16.05 126.63A16 16 0 0 0 32 144h88v56a32 32 0 0 0 64 0a8 8 0 0 0-16 0a16 16 0 0 1-32 0v-56h88a16 16 0 0 0 16-17.37ZM32 128a96.15 96.15 0 0 1 76.2-85.89C96.48 58 81.85 86.11 80.17 128H32Zm143.83 0c-1.68-41.89-16.31-70-28-85.94A96.07 96.07 0 0 1 224 128Z"/>`),
		g.Group(children),
	)
}