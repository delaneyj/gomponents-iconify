package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tachiyomisy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.576 34.542a8.811 8.811 0 0 0 7.391 3.327h4.435a7.413 7.413 0 0 0 7.391-7.392h0a7.413 7.413 0 0 0-7.39-7.391h-4.805a7.413 7.413 0 0 1-7.391-7.391h0a7.413 7.413 0 0 1 7.39-7.392h4.436c3.326 0 5.543.74 7.39 3.326"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.919 4.5L24.244 24L11.081 4.5m13.163 39V24"/>`),
		g.Group(children),
	)
}