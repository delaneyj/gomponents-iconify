package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LadderThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M211.76 214.63L160.26 73l10.54-29H184a4 4 0 0 0 0-8H88a4 4 0 0 0 0 8h10.29l-62 170.63a4 4 0 0 0 2.39 5.13A4.16 4.16 0 0 0 40 220a4 4 0 0 0 3.76-2.63L54.44 188h55.49l-9.69 26.63a4 4 0 0 0 2.39 5.13a4.12 4.12 0 0 0 1.37.24a4 4 0 0 0 3.76-2.63L127.17 164h57.66l19.41 53.37A4 4 0 0 0 208 220a4.16 4.16 0 0 0 1.37-.24a4 4 0 0 0 2.39-5.13ZM130.29 132H74.8l14.55-40h55.48Zm32-88l-14.55 40H92.26l14.54-40ZM57.35 180l14.54-40h55.49l-14.55 40Zm72.72-24L156 84.7l25.92 71.3Z"/>`),
		g.Group(children),
	)
}