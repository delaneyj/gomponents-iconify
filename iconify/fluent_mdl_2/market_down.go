package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarketDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 1024h128v640h-640v-128h421l-677-677l-384 384L3 477l90-90l675 674l384-384l768 768v-421z"/>`),
		g.Group(children),
	)
}