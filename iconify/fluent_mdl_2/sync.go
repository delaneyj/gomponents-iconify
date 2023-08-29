package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sync(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1024 1920q154 0 295-47t258-134t203-208t132-270l122 38q-50 167-149 304t-232 237t-294 153t-335 55q-137 0-267-34t-245-98t-214-157t-170-210v243H0v-512h512v128H196q59 117 146 211t196 161t231 103t255 37zM2048 256v512h-512V640h316q-59-117-146-211t-196-161t-231-103t-255-37q-154 0-295 47T471 309T268 517T136 787L14 749q50-166 149-304t232-237T689 55t335-55q137 0 267 34t245 98t214 157t170 210V256h128z"/>`),
		g.Group(children),
	)
}