package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func D(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M490 324V0h72v754h-72v-82c-53 58-132 96-216 96C122 768 0 648 0 498s122-270 274-270c84 0 163 38 216 96zM275 702c115 0 215-91 215-204S390 294 275 294c-114 0-207 91-207 204s93 204 207 204z"/>`),
		g.Group(children),
	)
}