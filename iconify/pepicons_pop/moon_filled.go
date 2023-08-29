package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoonFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M11.364 2.932a4.506 4.506 0 0 0-2.6 3.196a4.505 4.505 0 0 0 3.464 5.347a4.504 4.504 0 0 0 4.885-2.245c.489-.894 1.845-.57 1.877.449a9.045 9.045 0 0 1-.195 2.166c-1.035 4.87-5.815 7.98-10.678 6.947c-4.862-1.034-7.964-5.82-6.929-10.69c.974-4.58 5.283-7.644 9.895-7.078c1.008.124 1.21 1.498.28 1.908Z"/>`),
		g.Group(children),
	)
}