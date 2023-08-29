package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExercicesDeGainage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m5.5 26.828l27.943-5.858m-4.247.89v10.416h7.126m2.716-8.03a4.355 4.355 0 1 1-1.787-8.524a4.355 4.355 0 0 1 1.787 8.525ZM5.5 32.372V26.83"/>`),
		g.Group(children),
	)
}