package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SizeLegacy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 1152h128v768H0V128h1792v512h128v512zM128 1792h512V640h1024V256H128v1536zm640 0h512v-640h512V768H768v1024zm1152 0v-512h-512v512h512z"/>`),
		g.Group(children),
	)
}