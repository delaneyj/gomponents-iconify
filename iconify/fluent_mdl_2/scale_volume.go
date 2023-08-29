package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScaleVolume(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M512 128H219l338 339l-90 90l-339-338v293H0V0h512v128zM1536 0h512v512h-128V219l-339 338l-90-90l338-339h-293V0zM467 1491l90 90l-338 339h293v128H0v-512h128v293l339-338zm1453 338v-293h128v512h-512v-128h293l-338-339l90-90l339 338zM640 1408V640h768v768H640zm128-640v512h512V768H768z"/>`),
		g.Group(children),
	)
}