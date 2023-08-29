package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderOut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 13.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3zm0-2a.501.501 0 1 1-.002 1.002A.501.501 0 0 1 16 11.5zM20.5 3h-17a.5.5 0 0 0-.5.5v17a.5.5 0 0 0 .5.5h17a.5.5 0 0 0 .5-.5v-17a.5.5 0 0 0-.5-.5zM20 20H4V4h16v16zm-8-2.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3zm0-2a.501.501 0 1 1-.002 1.002A.501.501 0 0 1 12 15.5zm0-6a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3zm0-2a.501.501 0 1 1-.002 1.002A.501.501 0 0 1 12 7.5zm-4 6a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3zm0-2a.501.501 0 1 1-.002 1.002A.501.501 0 0 1 8 11.5zm4 2a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3zm0-2a.501.501 0 1 1-.002 1.002A.501.501 0 0 1 12 11.5z"/>`),
		g.Group(children),
	)
}