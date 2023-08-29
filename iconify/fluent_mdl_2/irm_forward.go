package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IRMForward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M128 960q0 115 30 221t84 198t130 169t168 130t199 84t221 30h320v71q-78 27-158 42t-162 15q-133 0-255-34t-228-97t-194-150t-149-195t-97-229T3 960q0-132 34-254t96-230t150-194t193-150t229-97T960 0q132 0 255 34t229 97t194 150t150 194t97 230t35 255q0 130-35 256l-107-107q14-76 14-149q0-115-30-221t-84-198t-130-169t-168-130t-199-84t-221-30q-115 0-221 30t-198 84t-169 130t-130 168t-84 199t-30 221zm256-64h1152v128H384V896zm1658 704l-317 317l-90-90l163-163h-518v-128h518l-163-163l90-90l317 317z"/>`),
		g.Group(children),
	)
}