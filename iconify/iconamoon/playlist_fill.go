package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaylistFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7 6a1 1 0 0 1-.5.866l-3 1.732A1 1 0 0 1 2 7.732V4.268a1 1 0 0 1 1.5-.866l3 1.732A1 1 0 0 1 7 6Zm-4 5a1 1 0 1 0 0 2h18a1 1 0 1 0 0-2H3Zm6-5a1 1 0 0 1 1-1h11a1 1 0 1 1 0 2H10a1 1 0 0 1-1-1ZM3 17a1 1 0 1 0 0 2h18a1 1 0 1 0 0-2H3Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}