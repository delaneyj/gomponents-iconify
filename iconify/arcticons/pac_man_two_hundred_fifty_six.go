package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PacManTwoHundredFiftySix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.363 25.565L3.5 27.764c4.689 10.109 15.484 18.245 28.843 15.765c10.775-2 12.941-14.895 6.86-24.067c-5.075-7.657-12.132-12.181-22.626-12.39c-5.42-.107-10.069 1.908-12.776 6.84Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.787 34.444c2.984-5.146 1.847-12.694-1.69-18.03C36.02 8.757 28.963 4.234 18.47 4.025c-5.42-.107-10.069 1.908-12.776 6.84M3.5 27.764l1.893-3.048M3.801 13.912l1.893-3.047m12.871 12.316L5.393 24.716m35.076 13.356l2.318-3.628"/>`),
		g.Group(children),
	)
}