package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ICursor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M832 128q-320 0-320 224v416h128v128H512v544q0 224 320 224h64v128h-64q-272 0-384-146q-112 146-384 146H0v-128h64q320 0 320-224V896H256V768h128V352q0-224-320-224H0V0h64q272 0 384 146Q560 0 832 0h64v128h-64z"/>`),
		g.Group(children),
	)
}