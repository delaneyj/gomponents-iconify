package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StatusCircleOuter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M960 128q132 0 254 34t230 97t194 150t150 195t97 229t35 255q0 132-34 254t-97 230t-150 194t-195 150t-229 97t-255 35q-132 0-254-34t-230-97t-194-150t-150-195t-97-229t-35-255q0-132 34-254t97-230t150-194t195-150t229-97t255-35z"/>`),
		g.Group(children),
	)
}