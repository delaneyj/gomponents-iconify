package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AsrLicence(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.959 43.55a21.484 21.484 0 1 1 10.59-10.586"/><circle cx="38" cy="38" r="7.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.222 34.811a12.997 12.997 0 1 1 3.577-3.57m2.633 2.902l-1.361 7.714h3.858"/>`),
		g.Group(children),
	)
}