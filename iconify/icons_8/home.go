package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Home(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m16 2.594l-.72.687l-13 13l1.44 1.44L5 16.437V28h9V18h4v10h9V16.437l1.28 1.282l1.44-1.44l-13-13l-.72-.686zm0 2.844l9 9V26h-5V16h-8v10H7V14.437l9-9z"/>`),
		g.Group(children),
	)
}