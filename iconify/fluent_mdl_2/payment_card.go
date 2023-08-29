package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaymentCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1888 256q33 0 62 12t51 35t34 51t13 62v1088q0 33-12 62t-35 51t-51 34t-62 13H160q-33 0-62-12t-51-35t-34-51t-13-62V416q0-33 12-62t35-51t51-34t62-13h1728zM160 384q-14 0-23 9t-9 23v224h1792V416q0-14-9-23t-23-9H160zm1728 1152q14 0 23-9t9-23V768H128v736q0 14 9 23t23 9h1728zm-480-384h256v128h-256v-128z"/>`),
		g.Group(children),
	)
}