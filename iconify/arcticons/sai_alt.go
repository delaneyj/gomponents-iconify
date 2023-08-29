package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SaiAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m23.914 43.5l6.957-7.64H16.956l6.958 7.64zm0-7.64V17.929L37.343 4.5m-17.098 9.761l-9.588-9.588"/>`),
		g.Group(children),
	)
}