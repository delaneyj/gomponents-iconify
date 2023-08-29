package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Move(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m245 1024l206 205l-91 91L0 960l360-360l91 91l-206 205h395v128H245zm1675-64l-356 355l-90-90l201-201h-395V896h395l-206-205l91-91l360 360zM695 446l-90-90L960 0l360 360l-91 91l-205-206v395H896V245L695 446zm534 1023l91 91l-360 360l-360-360l91-91l205 206v-395h128v395l205-206z"/>`),
		g.Group(children),
	)
}