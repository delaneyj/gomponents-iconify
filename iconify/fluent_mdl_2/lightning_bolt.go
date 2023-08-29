package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightningBolt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1223 768h660L603 2048H313l384-768H248L888 0h719l-384 768zM549 1920L1573 896h-557l384-768H967L455 1152h449l-384 768h29z"/>`),
		g.Group(children),
	)
}