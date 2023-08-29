package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stamp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 18h-6v-3.102a7.494 7.494 0 0 1 1.159-3.854a5.975 5.975 0 0 0-1.374-7.7a5.997 5.997 0 0 0-9.683 3.54a5.91 5.91 0 0 0 .748 4.174A7.257 7.257 0 0 1 12 14.794V18H6a2.002 2.002 0 0 0-2 2v4a2.002 2.002 0 0 0 2 2v2a2.002 2.002 0 0 0 2 2h16a2.002 2.002 0 0 0 2-2v-2a2.002 2.002 0 0 0 2-2v-4a2.002 2.002 0 0 0-2-2ZM12.07 7.244a4 4 0 1 1 7.373 2.773A9.76 9.76 0 0 0 18.066 14H13.94a9.543 9.543 0 0 0-1.376-3.974a3.937 3.937 0 0 1-.494-2.782ZM18 16v2h-4v-2Zm6 12H8v-2h16ZM6 24v-4h20v4Z"/>`),
		g.Group(children),
	)
}