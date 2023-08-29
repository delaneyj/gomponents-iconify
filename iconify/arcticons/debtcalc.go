package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Debtcalc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5 43.5h27a2 2 0 0 0 2-2v-35a2 2 0 0 0-2-2h-27a2 2 0 0 0-2 2v35a2 2 0 0 0 2 2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.533 7.963h24.935v6.575H11.533zm6.001 10.038v5.278h-6v-5.278zm-.001 8.327v5.278h-6v-5.278zm0 8.329v5.278h-6v-5.278zm9.468-16.656v5.278h-6v-5.278zM27 26.328v5.278h-6v-5.278zm0 8.329v5.278h-6v-5.278zm9.468-16.656v5.278h-6v-5.278zm-.001 8.327v5.278h-6v-5.278zm0 8.329v5.278h-6v-5.278zm-5.626-23.412h-2.028"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m32.996 10.407l1.13-.662v3"/>`),
		g.Group(children),
	)
}