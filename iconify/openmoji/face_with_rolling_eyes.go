package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceWithRollingEyes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<circle cx="36" cy="36" r="23" fill="#FCEA2B"/><circle cx="46" cy="32" r="6.5" fill="#FFF"/><circle cx="26" cy="32" r="6.5" fill="#FFF"/><circle cx="46.056" cy="28" r="2.556"/><circle cx="26.056" cy="28" r="2.556"/><circle cx="36" cy="36" r="23" fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"/><circle cx="46" cy="32" r="6.5" fill="none" stroke="#000" stroke-miterlimit="10" stroke-width="2"/><circle cx="26" cy="32" r="6.5" fill="none" stroke="#000" stroke-miterlimit="10" stroke-width="2"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M31 49.497h10"/>`),
		g.Group(children),
	)
}