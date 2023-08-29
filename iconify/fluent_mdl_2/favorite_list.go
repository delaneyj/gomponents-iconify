package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FavoriteList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 768v128h-878q-36-118-72-233t-74-234q-37 118-73 233t-73 234H376q102 78 202 156t203 155q-40 124-78 246t-76 247l397-306v162l-640 492l248-794L0 768h784L1024 0l240 768h784zm-896 384h896v128h-896v-128zm0 384h896v128h-896v-128z"/>`),
		g.Group(children),
	)
}