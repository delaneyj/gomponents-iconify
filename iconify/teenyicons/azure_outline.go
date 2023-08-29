package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AzureOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="m3.5 4.5l-3 7h3l4-11l-4 4Zm11 9l-5-11l-2 5l3 4l-6 2h10Z"/>`),
		g.Group(children),
	)
}