package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Droplets(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.5 28C8.5 14.23 24 4.5 24 4.5S39.5 14.3 39.5 28a15.5 15.5 0 0 1-31 0Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19 28a5 5 0 0 0 10 0Z"/><circle cx="20.098" cy="24.071" r=".75" fill="currentColor"/><circle cx="27.929" cy="24.071" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}