package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OEM(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1600 1024h192v896H128V256h896v192L1472 0l576 576l-448 448zm-546-448l418 418l418-418l-418-418l-418 418zm-30 128v320h320l-320-320zm-768 320h640V384H256v640zm640 128H256v640h640v-640zm128 0v640h640v-640h-640z"/>`),
		g.Group(children),
	)
}