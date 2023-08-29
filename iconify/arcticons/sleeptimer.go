package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sleeptimer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="24" r="7" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m21.89 19.971l3.23.581l-4.09 4.393l3.23.581m3.029-6.258h5.3l-5.3 8h5.3"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m29.495 13.793l12.112-2.131l-8.975 19.88l12.112-2.13"/>`),
		g.Group(children),
	)
}