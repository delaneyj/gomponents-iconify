package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Babbel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".866" d="M36.456 24a4.823 4.823 0 0 1 4.868 5.91a7.356 7.356 0 0 1-6.952 5.91h-9.751l4.168-23.639h9.751a4.823 4.823 0 0 1 4.868 5.91A7.356 7.356 0 0 1 36.456 24Zm0 0h-9.751m-6.495 0H4.5m8.661-4.568l.579-3.287m-2.77 15.71L12.355 24"/>`),
		g.Group(children),
	)
}