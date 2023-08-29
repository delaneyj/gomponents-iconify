package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1760 0q66 0 113 47t47 113v1216q0 66-47 113t-113 47H160q-66 0-113-47T0 1376V160Q0 94 47 47T160 0h1600zM160 128q-13 0-22.5 9.5T128 160v224h1664V160q0-13-9.5-22.5T1760 128H160zm1600 1280q13 0 22.5-9.5t9.5-22.5V768H128v608q0 13 9.5 22.5t22.5 9.5h1600zM256 1280v-128h256v128H256zm384 0v-128h384v128H640z"/>`),
		g.Group(children),
	)
}