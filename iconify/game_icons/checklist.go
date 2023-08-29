package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Checklist(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m122.31 84.615l-2.85 8.54l-11.394 34.185l-5.703-5.703L96 115.27L83.27 128l6.367 6.363l26.297 26.297l20.605-61.814l2.845-8.537l-17.076-5.695zM151 119v18h242v-18H151zm0 64v18h242v-18H151zm0 64v18h242v-18H151zm-28.69 29.615l-2.85 8.54l-11.394 34.185l-5.703-5.703L96 307.27L83.27 320l6.367 6.363l26.297 26.297l20.605-61.814l2.845-8.537l-17.076-5.695zM151 311v18h242v-18H151zm0 64v18h242v-18H151z"/>`),
		g.Group(children),
	)
}