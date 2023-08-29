package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Coini(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.49 8.73C-9.25 8.55 19.36 59.79 38.29 27M15.59 5.5l19.65 34.05M10.05 8.45L29.71 42.5"/>`),
		g.Group(children),
	)
}