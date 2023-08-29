package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chubbyclick(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.45 9.39H4.5m0 16.04h16.15M4.5 17.41h24.95"/><circle cx="31.05" cy="32.02" r="6.59" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.64 32.02V9.39h5.86"/>`),
		g.Group(children),
	)
}