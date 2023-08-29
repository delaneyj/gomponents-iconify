package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IRMForwardMirrored(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M960 0q132 0 255 34t229 97t194 150t150 194t97 230t35 255q0 132-34 255t-97 229t-150 194t-194 150t-230 97t-255 35q-95 0-192-20v-108h192q115 0 221-30t198-84t169-130t130-168t84-199t30-221q0-115-30-221t-84-198t-130-169t-168-130t-199-84t-221-30q-115 0-221 30t-198 84t-169 130t-130 168t-84 199t-30 221q0 64 9 127t30 125L67 1312q-33-85-48-173T3 960q0-132 34-254t96-230t150-194t193-150t229-97T960 0zm576 1024H384V896h1152v128zM250 1664l163 163l-90 90L6 1600l317-317l90 90l-163 163h518v128H250z"/>`),
		g.Group(children),
	)
}