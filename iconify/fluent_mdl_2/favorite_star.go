package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FavoriteStar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1416 1254l248 794l-640-492l-640 492l248-794L0 768h784L1024 0l240 768h784l-632 486zm5 446q-38-124-76-246t-78-247q103-77 203-155t202-156h-502l-146-467l-146 467H376q102 78 202 156t203 155q-40 124-78 246t-76 247l397-306l397 306z"/>`),
		g.Group(children),
	)
}