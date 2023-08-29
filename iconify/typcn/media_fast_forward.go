package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MediaFastForward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.053 6.912A1.797 1.797 0 0 0 12 8.201v9a1.8 1.8 0 0 0 3.053 1.287C17.434 16.174 21 12.701 21 12.701s-3.566-3.474-5.947-5.789zm-9 0A1.797 1.797 0 0 0 3 8.201v9a1.8 1.8 0 0 0 3.053 1.287C8.434 16.174 12 12.701 12 12.701S8.434 9.227 6.053 6.912z"/>`),
		g.Group(children),
	)
}