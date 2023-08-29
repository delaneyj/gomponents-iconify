package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DRM(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M960 0q132 0 255 34t229 97t194 150t150 194t97 230t35 255q0 132-34 255t-97 229t-150 194t-194 150t-230 97t-255 35q-132 0-255-34t-229-97t-194-150t-150-194t-97-229T0 960q0-132 34-255t97-229t150-194t194-150t229-97T960 0zm576 896H384v128h1152V896z"/>`),
		g.Group(children),
	)
}