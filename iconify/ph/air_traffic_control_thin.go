package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirTrafficControlThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M225.83 65.11A12 12 0 0 0 216 60h-84V20h20a4 4 0 0 0 0-8h-48a4 4 0 0 0 0 8h20v40H40a12 12 0 0 0-11.28 16.1l26.19 72a12 12 0 0 0 11.27 7.9H100v68a4 4 0 0 0 8 0v-68h40v68a4 4 0 0 0 8 0v-68h33.82a12 12 0 0 0 11.27-7.9l26.19-72a12 12 0 0 0-1.45-10.99ZM107.34 148L92.79 68h70.42l-14.55 80Zm-44.92-2.63l-26.18-72A4 4 0 0 1 40 68h44.66l14.54 80h-33a4 4 0 0 1-3.78-2.63Zm157.34-72l-26.18 72a4 4 0 0 1-3.76 2.63h-33l14.55-80H216a4 4 0 0 1 3.76 5.37Z"/>`),
		g.Group(children),
	)
}