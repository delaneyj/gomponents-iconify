package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zoom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 1978q0 28-17 49t-47 21q-26 0-45-19l-785-784q-96 79-210 121t-240 42q-97 0-187-25t-168-71t-142-110t-111-143t-71-168T0 704q0-97 25-187t71-168t110-142T349 96t168-71T704 0q97 0 187 25t168 71t142 110t111 143t71 168t25 187q0 125-42 239t-121 211l57 57q42 42 102 100t130 127t142 141t139 140t119 123t83 91t31 45zM704 1280q119 0 224-45t183-124t123-183t46-224q0-119-45-224t-124-183t-183-123t-224-46q-119 0-224 45T297 297T174 480t-46 224q0 119 45 224t124 183t183 123t224 46z"/>`),
		g.Group(children),
	)
}