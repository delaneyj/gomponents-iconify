package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarFull(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10 16.207l-6.173 3.246l1.179-6.874L.01 7.71l6.902-1.003L10 .453l3.087 6.254l6.902 1.003l-4.995 4.869l1.18 6.874z"/>`),
		g.Group(children),
	)
}