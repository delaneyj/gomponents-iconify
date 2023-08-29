package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kinemaster(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.374 13.533v20.934m11.252 0L20.991 24l8.635-10.467M20.991 24h-2.617"/><circle cx="24" cy="6" r=".75" fill="currentColor"/><circle cx="41.119" cy="18.438" r=".75" fill="currentColor"/><circle cx="34.58" cy="38.562" r=".75" fill="currentColor"/><circle cx="13.42" cy="38.562" r=".75" fill="currentColor"/><circle cx="6.881" cy="18.438" r=".75" fill="currentColor"/><circle cx="24" cy="6" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}