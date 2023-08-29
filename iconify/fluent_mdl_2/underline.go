package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Underline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1011 1792q-161 0-274-46t-184-133t-105-209t-33-275V128h192v988q0 109 21 201t71 159t131 106t201 38q115 0 193-36t127-101t69-154t21-195V128h192v975q0 158-34 285t-109 217t-193 138t-286 49zm-627 128h1280v128H384v-128z"/>`),
		g.Group(children),
	)
}