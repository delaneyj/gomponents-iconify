package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronCircleLeftThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M30 16c0 7.732-6.268 14-14 14S2 23.732 2 16S8.268 2 16 2s14 6.268 14 14Zm-12.707 6.707a1 1 0 0 0 1.414-1.414L13.414 16l5.293-5.293a1 1 0 0 0-1.414-1.414l-5.879 5.878l-.028.033a1.662 1.662 0 0 1-.073.08a.997.997 0 0 0-.293.716c-.003.26.097.484.293.718c.041.05.05.06.1.11h.002l5.878 5.88Z"/>`),
		g.Group(children),
	)
}