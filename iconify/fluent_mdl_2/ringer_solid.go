package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RingerSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1792 1536v128h-512q0 53-20 99t-55 82t-81 55t-100 20q-53 0-99-20t-82-55t-55-81t-20-100H256v-128h128V768q0-88 23-170t64-153t100-129t130-100t153-65t170-23q88 0 170 23t153 64t129 100t100 130t65 153t23 170v768h128z"/>`),
		g.Group(children),
	)
}