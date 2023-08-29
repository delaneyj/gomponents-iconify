package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AhsParking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 17.833h12.333v12.333H5.5zm24.667 0H42.5v12.333H30.167zM17.833 30.167h12.333V42.5H17.833zm9.334 0l-9.334-9.334V5.5h12.334v24.667h-3z"/>`),
		g.Group(children),
	)
}