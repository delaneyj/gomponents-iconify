package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Qrcode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M384 1024v128H256v-128h128zm0-768v128H256V256h128zm768 0v128h-128V256h128zM128 1279h384V896H128v383zm0-767h384V128H128v384zm768 0h384V128H896v384zM640 768v640H0V768h640zm512 512v128h-128v-128h128zm256 0v128h-128v-128h128zm0-512v384h-384v-128H896v384H768V768h384v128h128V768h128zM640 0v640H0V0h640zm768 0v640H768V0h640z"/>`),
		g.Group(children),
	)
}