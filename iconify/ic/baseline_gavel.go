package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BaselineGavel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m5.25 8.069l2.83-2.827l14.134 14.15l-2.83 2.827zm4.236-4.242L12.314.998l5.657 5.656l-2.828 2.83zM.999 12.315l2.828-2.829l5.657 5.657l-2.828 2.828zM1 21h12v2H1z"/>`),
		g.Group(children),
	)
}