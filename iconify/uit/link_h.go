package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkH(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.5 16H7a4 4 0 1 1 0-8h3.5a.5.5 0 0 0 0-1H7a5 5 0 0 0 0 10h3.5a.5.5 0 0 0 0-1zM8 12a.5.5 0 0 0 .5.5h7a.5.5 0 0 0 0-1h-7a.5.5 0 0 0-.5.5zm9-5h-3.5a.5.5 0 0 0 0 1H17a4 4 0 1 1 0 8h-3.5a.5.5 0 0 0 0 1H17a5 5 0 0 0 0-10z"/>`),
		g.Group(children),
	)
}