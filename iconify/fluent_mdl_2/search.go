package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Search(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1344 0q97 0 187 25t168 71t142 110t111 143t71 168t25 187q0 97-25 187t-71 168t-110 142t-143 111t-168 71t-187 25q-125 0-239-42t-211-121l-785 784q-19 19-45 19t-45-19t-19-45q0-26 19-45l784-785q-79-96-121-210t-42-240q0-97 25-187t71-168t110-142T989 96t168-71t187-25zm0 1280q119 0 224-45t183-124t123-183t46-224q0-119-45-224t-124-183t-183-123t-224-46q-119 0-224 45T937 297T814 480t-46 224q0 119 45 224t124 183t183 123t224 46z"/>`),
		g.Group(children),
	)
}