package charm

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 10.25h9.5m-8.5-4.5h9.5m-2.5-4l-1.5 12.5m-2.5-12.5l-1.5 12.5"/>`),
		g.Group(children),
	)
}