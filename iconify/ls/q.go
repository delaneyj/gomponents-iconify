package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Q(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M490 210v-97h72v754h-72V558c-53 58-132 96-216 96C122 654 0 534 0 384s122-271 274-271c84 0 163 39 216 97zM275 587c115 0 215-90 215-203S390 180 275 180c-114 0-207 91-207 204s93 203 207 203z"/>`),
		g.Group(children),
	)
}