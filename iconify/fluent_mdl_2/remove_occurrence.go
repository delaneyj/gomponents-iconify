package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RemoveOccurrence(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 256v512h-512V640h316q-59-117-146-211t-196-161t-231-103t-255-37q-154 0-295 47T471 309T268 517T136 787L14 749q50-166 149-304t232-237T689 55t335-55q137 0 267 34t245 98t214 157t170 210V256h128zM1024 1920q39 0 77-4t77-13l108 108q-31 8-62 14t-64 12l-8-64l8 64q-34 5-67 8t-69 3q-138 0-267-34t-245-98t-213-157t-171-210v243H0v-512h512v128H196q59 117 146 211t196 161t231 103t255 37zm1021-418l-226 226l226 227l-90 90l-227-226l-227 227l-90-91l227-227l-227-227l90-90l227 227l227-227l90 91z"/>`),
		g.Group(children),
	)
}