package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MuvizEdge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.362 15.424L16.879 5.5L6.319 42.304l29.602-18.778l-5.062-4.791"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.557 16.032L22.638 5.696L12.079 42.5l9.041-5.735m5.407-3.43l15.154-9.613l-5.398-5.109"/>`),
		g.Group(children),
	)
}