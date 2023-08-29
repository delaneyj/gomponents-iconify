package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronCircleUpThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16 30C8.268 30 2 23.732 2 16S8.268 2 16 2s14 6.268 14 14s-6.268 14-14 14ZM9.293 17.293a1 1 0 1 0 1.414 1.414L16 13.414l5.293 5.293a1 1 0 0 0 1.414-1.414l-5.878-5.879l-.033-.028a1.662 1.662 0 0 1-.08-.073a.997.997 0 0 0-.716-.293c-.26-.003-.484.097-.718.293c-.05.041-.06.05-.11.1v.002l-5.88 5.878Z"/>`),
		g.Group(children),
	)
}