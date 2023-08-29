package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimesRectangleO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="m1257 983l-146 146q-10 10-23 10t-23-10L896 960l-169 169q-10 10-23 10t-23-10L535 983q-10-10-10-23t10-23l169-169l-169-169q-10-10-10-23t10-23l146-146q10-10 23-10t23 10l169 169l169-169q10-10 23-10t23 10l146 146q10 10 10 23t-10 23l-169 169l169 169q10 10 10 23t-10 23zM256 1280h1280V256H256v1024zM1792 160v1216q0 66-47 113t-113 47H160q-66 0-113-47T0 1376V160Q0 94 47 47T160 0h1472q66 0 113 47t47 113z"/>`),
		g.Group(children),
	)
}