package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SyncToPC(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 256v512h-512V640h316q-59-117-146-211t-196-161t-231-103t-255-37q-154 0-295 47T471 309T268 517T136 787L14 749q50-166 149-304t232-237T689 55t335-55q137 0 267 34t245 98t214 157t170 210V256h128zM196 1408q59 117 146 211t196 161t231 103t255 37h128v118q-32 4-64 7t-64 3q-136 0-266-34t-246-100t-214-157t-170-208v243H0v-512h511v128H196zm956-256h896v640h-384v128h127v128h-383v-128h128v-128h-384v-640zm128 512h640v-384h-640v384z"/>`),
		g.Group(children),
	)
}