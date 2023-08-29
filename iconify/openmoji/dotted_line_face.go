package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DottedLineFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<circle cx="35.977" cy="35.958" r="23" fill="none" stroke="#000" stroke-dasharray="5.16 5.16" stroke-linecap="round" stroke-miterlimit="10" stroke-width="2"/><circle cx="35.977" cy="35.958" r="23" fill="none" stroke="#f4aa41" stroke-dasharray="5.16 5.16" stroke-linecap="round" stroke-miterlimit="10" stroke-width="2"/><path d="M29.977 32.958a3 3 0 1 1-3-3a3.001 3.001 0 0 1 3 3m18 0a3 3 0 1 1-3-3a3.001 3.001 0 0 1 3 3"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-miterlimit="10" stroke-width="2" d="M29.977 43.796h12"/>`),
		g.Group(children),
	)
}