package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circles(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="25.193" cy="24" r="10.574" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="21.331" cy="12.115" r="8.613" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="35.303" cy="16.655" r="8.613" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="35.303" cy="31.345" r="8.613" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="21.331" cy="35.885" r="8.613" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="12.697" cy="24" r="8.613" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}