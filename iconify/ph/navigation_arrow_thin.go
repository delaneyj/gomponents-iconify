package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NavigationArrowThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M228 102L51.93 36.67A12 12 0 0 0 36.69 52L102 228a11.81 11.81 0 0 0 11.31 8h.22a11.82 11.82 0 0 0 11.26-8.47L148.32 151a4 4 0 0 1 2.65-2.65l76.56-23.55A12 12 0 0 0 228 102Zm-2.83 15.13l-76.57 23.56a12 12 0 0 0-7.94 7.94l-23.55 76.56a3.89 3.89 0 0 1-3.76 2.82a3.93 3.93 0 0 1-3.85-2.69v-.08L44.22 49.32a3.93 3.93 0 0 1 1-4.14A4 4 0 0 1 48 44a3.86 3.86 0 0 1 1.25.21l176.08 65.32a4 4 0 0 1-.13 7.6Z"/>`),
		g.Group(children),
	)
}