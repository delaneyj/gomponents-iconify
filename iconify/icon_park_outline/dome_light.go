package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DomeLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24.778 8c-11.046 0-20 8.954-20 20h40c0-11.046-8.954-20-20-20Zm0-4v4m0 30c-5.523 0-10-4.477-10-10h20c0 5.523-4.477 10-10 10Zm16.034.977l-2.068-2.954m-27.691.202l-2.55 2.55M34.778 42l-1.147-1.638m-17.72.112l-1.414 1.414"/>`),
		g.Group(children),
	)
}