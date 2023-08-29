package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Print(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M363 107q26 0 45 18.5t19 45.5v128h-86v85H85v-85H0V171q0-27 18.5-45.5T64 107h299zm-64 234V235H128v106h171zm63.5-149q8.5 0 15-6.5t6.5-15t-6.5-15t-15-6.5t-15 6.5t-6.5 15t6.5 15t15 6.5zM341 0v85H85V0h256z"/>`),
		g.Group(children),
	)
}