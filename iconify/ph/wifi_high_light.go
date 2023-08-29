package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiHighLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M138 204a10 10 0 1 1-10-10a10 10 0 0 1 10 10Zm97.81-115.45a170 170 0 0 0-215.62 0a6 6 0 1 0 7.62 9.27a158 158 0 0 1 200.38 0a6 6 0 1 0 7.62-9.27Zm-32.08 35.79a122 122 0 0 0-151.46 0a6 6 0 0 0 7.46 9.41a110 110 0 0 1 136.54 0A6 6 0 0 0 200 135a6 6 0 0 0 3.73-10.7Zm-32.2 35.81a74 74 0 0 0-87.06 0a6 6 0 0 0 7.06 9.7a62 62 0 0 1 72.94 0a6 6 0 0 0 8.38-1.32a6 6 0 0 0-1.32-8.38Z"/>`),
		g.Group(children),
	)
}