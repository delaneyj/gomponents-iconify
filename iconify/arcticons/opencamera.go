package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Opencamera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.19 6.42h11.63l3.86 3.83h5.75a4.06 4.06 0 0 1 4.07 4.07v23.19a4.05 4.05 0 0 1-4.07 4.07H8.57a4.06 4.06 0 0 1-4.07-4.07V14.32a4.07 4.07 0 0 1 4.07-4.07h5.76l3.86-3.83Zm5.74 9.79a9.71 9.71 0 1 0 9.77 9.71v-.1a9.71 9.71 0 0 0-9.77-9.61Zm0 3.92a5.79 5.79 0 0 1 5.83 5.73v.06A5.79 5.79 0 1 1 24 20.13Z"/>`),
		g.Group(children),
	)
}