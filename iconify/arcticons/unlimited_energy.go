package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnlimitedEnergy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="14.94" cy="29.388" r="2.729" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="33.408" cy="29.231" r="2.729" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24.024" cy="17.512" r="2.729" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24.024" cy="40.771" r="2.729" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.024 20.24v17.803M13.376 27.152C4.562 14.07 15.906 4.5 24.116 4.5s19.026 9.828 10.72 22.406"/>`),
		g.Group(children),
	)
}