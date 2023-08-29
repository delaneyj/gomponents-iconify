package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxLightDoubleVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M12 22V0h2v22h-2m-4 0V0h2v22H8Z"/>`),
		g.Group(children),
	)
}