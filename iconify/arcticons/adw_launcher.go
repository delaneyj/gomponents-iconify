package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdwLauncher(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 27.506L24 8.006l-19.5 19.5m22.775 12.021L24 36.252l-3.275 3.275"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.195 29.119c2.88 6.41 9.321 10.875 16.805 10.875S37.925 35.53 40.805 29.12m-8.646-12.955v-4.563h4.92v9.483"/>`),
		g.Group(children),
	)
}