package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Leak(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M64 0q0 27-18.5 45.5T0 64V0h64zm171 0q0 97-69 166T0 235v-43q80 0 136-56T192 0h43zm-86 0q0 62-43.5 105.5T0 149v-42q44 0 75.5-31.5T107 0h42zm0 384q0-97 69-166t166-69v43q-80 0-136 56t-56 136h-43zm171 0q0-27 18.5-45.5T384 320v64h-64zm-85 0q0-62 43.5-105.5T384 235v42q-44 0-75.5 31.5T277 384h-42z"/>`),
		g.Group(children),
	)
}