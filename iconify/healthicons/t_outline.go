package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13 12a3 3 0 0 1 3-3h16a3 3 0 1 1 0 6h-5v21a3 3 0 1 1-6 0V15h-5a3 3 0 0 1-3-3Zm3-1a1 1 0 1 0 0 2h7v23a1 1 0 1 0 2 0V13h7a1 1 0 1 0 0-2H16Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}