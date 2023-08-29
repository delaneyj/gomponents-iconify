package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Forerkort(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 37.5v-27a2 2 0 0 0-2-2h-35a2 2 0 0 0-2 2v27a2 2 0 0 0 2 2h35a2 2 0 0 0 2-2Zm-19.173-20h14.77m-14.77 6.252h14.77m-14.77 6.253h14.77"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.558 30.005H9.14s-.276-5.49 5.21-5.49c5.461 0 5.21 5.49 5.21 5.49Z"/><circle cx="14.349" cy="20.069" r="2.57" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}