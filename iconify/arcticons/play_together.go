package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayTogether(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m16.037 20.943l19.417-6.479a4.688 4.688 0 0 1 6.11 3.777c1.003 6.58.71 17.406-10.455 22.081a24.264 24.264 0 0 1-14.582 1.722"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m31.963 27.057l-19.417 6.479a4.688 4.688 0 0 1-6.11-3.777c-1.003-6.58-.71-17.406 10.454-22.081a24.264 24.264 0 0 1 14.583-1.722"/>`),
		g.Group(children),
	)
}