package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatHeartFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.455 19L2 22.5V4a1 1 0 0 1 1-1h18a1 1 0 0 1 1 1v14a1 1 0 0 1-1 1H6.455Zm5.563-4.3l3.358-3.359a2.25 2.25 0 1 0-3.181-3.182l-.177.177l-.177-.177a2.25 2.25 0 0 0-3.182 3.182l3.359 3.359Z"/>`),
		g.Group(children),
	)
}