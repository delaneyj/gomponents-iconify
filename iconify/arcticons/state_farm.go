package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StateFarm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<ellipse cx="14.25" cy="29.686" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="7.06" ry="4.035"/><ellipse cx="14.25" cy="29.633" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="9.75" ry="6.548"/><ellipse cx="24" cy="18.352" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="7.06" ry="4.035"/><ellipse cx="24" cy="18.299" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="9.75" ry="6.548"/><ellipse cx="33.75" cy="29.754" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="7.06" ry="4.035"/><ellipse cx="33.75" cy="29.701" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="9.75" ry="6.548"/>`),
		g.Group(children),
	)
}