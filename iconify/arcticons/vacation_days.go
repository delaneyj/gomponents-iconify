package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VacationDays(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.287 42.03a21.54 21.54 0 0 1 30.167-30.698m0 0L11.287 42.031"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.662 33.508C2.642 10.55 12.786 2.656 33.38 19.548m-3.728 9.833L42.5 42.201"/>`),
		g.Group(children),
	)
}