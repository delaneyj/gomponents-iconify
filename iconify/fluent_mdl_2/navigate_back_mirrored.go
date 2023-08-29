package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NavigateBackMirrored(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M960 0Q827 0 705 34t-230 96t-194 150t-150 195t-97 229T0 960q0 133 34 255t96 230t150 194t195 150t229 97t256 34q133 0 255-34t230-96t194-150t150-195t97-229t34-256q0-133-34-255t-96-230t-150-194t-195-150t-229-97T960 0zm0 1792q-115 0-221-30t-198-84t-169-130t-130-168t-84-199t-30-221q0-115 30-221t84-198t130-169t168-130t199-84t221-30q115 0 221 30t198 84t169 130t130 168t84 199t30 221q0 115-30 221t-84 198t-130 169t-168 130t-199 84t-221 30zm233-896H512v128h681l-278 274l90 92l434-430l-434-430l-90 92l278 274z"/>`),
		g.Group(children),
	)
}