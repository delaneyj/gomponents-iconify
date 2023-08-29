package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DivideOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 8.5c-1.654 0-3-1.346-3-3s1.346-3 3-3s3 1.346 3 3s-1.346 3-3 3zm0-4a1.001 1.001 0 0 0 0 2a1.001 1.001 0 0 0 0-2zm0 17c-1.654 0-3-1.346-3-3s1.346-3 3-3s3 1.346 3 3s-1.346 3-3 3zm0-4a1.001 1.001 0 0 0 0 2a1.001 1.001 0 0 0 0-2zm6-2.5H6c-1.654 0-3-1.346-3-3s1.346-3 3-3h12c1.654 0 3 1.346 3 3s-1.346 3-3 3zM6 11a1.001 1.001 0 0 0 0 2h12a1.001 1.001 0 0 0 0-2H6z"/>`),
		g.Group(children),
	)
}