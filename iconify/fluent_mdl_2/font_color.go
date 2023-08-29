package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FontColor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M256 1920h1408v128H256v-128zm387-640l-170 512H338L893 128h135l554 1664h-135l-171-512H643zm317-949l-274 821h548L960 331z"/>`),
		g.Group(children),
	)
}