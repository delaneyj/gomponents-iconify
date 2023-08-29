package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandDatabricks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 17l9 5l9-5v-3l-9 5l-9-5v-3l9 5l9-5V8l-9 5l-9-5l9-5l5.418 3.01"/>`),
		g.Group(children),
	)
}