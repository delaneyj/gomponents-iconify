package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PingUtils(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.829 26.801h9.202v9.202h-9.202zm0-14.804h9.202v9.202h-9.202zm14.389 4.601h9.954m-9.954 14.804h9.954"/><rect width="37.057" height="37.057" x="5.536" y="5.583" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" ry="2"/>`),
		g.Group(children),
	)
}