package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListUiAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.5 6a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3zm0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1zm4 0h14a.5.5 0 0 0 0-1h-14a.5.5 0 0 0 0 1zm14 4h-10a.5.5 0 0 0 0 1h10a.5.5 0 0 0 0-1zm0 5h-6a.5.5 0 0 0 0 1h6a.5.5 0 0 0 0-1zm-14-6a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3zm0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1zm4 3a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3zm0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1z"/>`),
		g.Group(children),
	)
}