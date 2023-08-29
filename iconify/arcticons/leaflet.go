package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Leaflet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 4.5s-11.069 9.052-13.722 22.165S21.537 43.5 24 43.5M10.506 25.56l10.93 6.512"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 21.366l-2.565 4.885v16.776m-6.88-27.177L24 21.366M24 4.5s11.069 9.052 13.722 22.165S26.463 43.5 24 43.5m13.494-17.94l-10.93 6.512"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 21.366l2.565 4.885v16.776m6.88-27.177L24 21.366m0 0V4.5"/>`),
		g.Group(children),
	)
}