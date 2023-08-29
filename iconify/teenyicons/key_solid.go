package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeySolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7 4a4 4 0 1 1 1.547 3.16l-3.34 3.34l1.647 1.646l-.708.708L4.5 11.207L3.207 12.5l1.647 1.646l-.708.708L2.5 13.207L.854 14.854l-.708-.708L7.84 6.453A3.983 3.983 0 0 1 7 4Zm4-3a3 3 0 1 0 0 6a3 3 0 0 0 0-6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}