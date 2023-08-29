package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpLeftTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M13.707 12.293L5.414 4H9a1 1 0 0 0 0-2H3a1.002 1.002 0 0 0-1 1h-.001v6a1 1 0 0 0 2 0V5.414l8.293 8.293a.997.997 0 0 0 1.414 0a.999.999 0 0 0 0-1.414z"/>`),
		g.Group(children),
	)
}