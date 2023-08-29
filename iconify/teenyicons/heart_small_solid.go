package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartSmallSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M6.765 5.235a1.79 1.79 0 1 0-2.53 2.53L7.5 11.03l3.265-3.265a1.79 1.79 0 0 0-2.53-2.53L7.5 5.97l-.735-.735Z"/>`),
		g.Group(children),
	)
}