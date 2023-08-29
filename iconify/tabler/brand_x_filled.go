package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandXFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M8.267 3a1 1 0 0 1 .73.317l.076.092l4.274 5.828l5.946-5.944a1 1 0 0 1 1.497 1.32l-.083.094l-6.163 6.162l6.262 8.54a1 1 0 0 1-.697 1.585L20 21h-4.267a1 1 0 0 1-.73-.317l-.076-.092l-4.276-5.829l-5.944 5.945a1 1 0 0 1-1.497-1.32l.083-.094l6.161-6.163l-6.26-8.539a1 1 0 0 1 .697-1.585L4 3h4.267z"/></g>`),
		g.Group(children),
	)
}