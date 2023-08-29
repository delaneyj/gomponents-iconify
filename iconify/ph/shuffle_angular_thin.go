package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShuffleAngularThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M234.83 181.17a4 4 0 0 1 0 5.66l-24 24a4 4 0 0 1-5.66-5.66L222.34 188h-50.22a12 12 0 0 1-9.77-5L87.14 77.67A4 4 0 0 0 83.88 76H32a4 4 0 0 1 0-8h51.88a12 12 0 0 1 9.77 5l75.21 105.31a4 4 0 0 0 3.26 1.67h50.22l-17.17-17.17a4 4 0 0 1 5.66-5.66Zm-89.49-77.44a4 4 0 0 0 5.58-.93l17.94-25.13a4 4 0 0 1 3.26-1.67h50.22l-17.17 17.17a4 4 0 0 0 5.66 5.66l24-24a4 4 0 0 0 0-5.66l-24-24a4 4 0 0 0-5.66 5.66L222.34 68h-50.22a12 12 0 0 0-9.77 5l-17.94 25.15a4 4 0 0 0 .93 5.58Zm-34.68 48.54a4 4 0 0 0-5.58.93l-17.94 25.13a4 4 0 0 1-3.26 1.67H32a4 4 0 0 0 0 8h51.88a12 12 0 0 0 9.77-5l17.94-25.13a4 4 0 0 0-.93-5.6Z"/>`),
		g.Group(children),
	)
}