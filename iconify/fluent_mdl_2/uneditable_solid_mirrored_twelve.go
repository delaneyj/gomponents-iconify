package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UneditableSolidMirroredTwelve(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M574 994L292 712l422-422l280 284l-420 420zm1474 1052l-753-333l-420-419l-729 729l-121-121L1902 25l121 121l-730 730l407 411l348 759zM140 562q-32-31-58-60t-44-62t-28-70t-10-84q0-61 23-113t63-91t93-60T294 0q47 0 83 10t68 28t59 44t60 57L140 562z"/>`),
		g.Group(children),
	)
}