package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func File(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M400.86 0L174.791 218.249v73.93h302.582V0H400.86zm179.219 0v391.27H174.791V1200h850.418V0h-445.13z"/>`),
		g.Group(children),
	)
}