package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandVlc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m13.79 4.337l3.101 9.305c.33.985-.113 2.07-1.02 2.499a9.148 9.148 0 0 1-7.742 0c-.907-.428-1.35-1.514-1.02-2.499l3.1-9.305C10.476 3.537 11.194 3 12 3c.807 0 1.525.537 1.79 1.337z"/><path d="M7 14H5.571a2 2 0 0 0-1.923 1.45l-.571 2A2 2 0 0 0 5 20h13.998a2 2 0 0 0 1.923-2.55l-.572-2A2 2 0 0 0 18.426 14H17"/></g>`),
		g.Group(children),
	)
}