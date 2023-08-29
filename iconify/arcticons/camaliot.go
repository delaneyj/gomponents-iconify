package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Camaliot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="14.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.5 24h29m-27.023-7.307h25.046M11.478 31.307h25.044M24 38.5v-29"/><ellipse cx="24" cy="24" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="7.307" ry="14.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.619 11.675a21.495 21.495 0 1 1-5.362-5.341"/><circle cx="39.203" cy="8.797" r="3.758" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}