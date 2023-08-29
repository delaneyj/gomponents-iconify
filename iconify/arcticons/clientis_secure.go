package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClientisSecure(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.331 4.5C19.428 9.689 32.233 21.49 41.69 24M11.332 7.896c2.919 4.892 14.464 15.658 23.7 18.98M6.31 10.742c2.96 4.96 14.788 15.959 24.085 19.114"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.331 43.5C19.428 38.311 32.233 26.51 41.69 24M11.332 40.104c2.919-4.891 14.464-15.658 23.7-18.98M6.31 37.258c2.96-4.96 14.789-15.959 24.086-19.114"/>`),
		g.Group(children),
	)
}