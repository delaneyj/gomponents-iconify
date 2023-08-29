package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarromDiscPool(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="18.983" r="14.482" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.449 29.009c2.498 1.003 4.034 2.365 4.034 3.865c0 3.083-6.484 5.582-14.483 5.582S9.518 35.957 9.518 32.874c0-1.5 1.535-2.862 4.034-3.865"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.518 32.874v5.044C9.518 41 16 43.5 24 43.5s14.482-2.5 14.482-5.582v-5.044"/>`),
		g.Group(children),
	)
}