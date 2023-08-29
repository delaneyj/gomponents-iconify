package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThermometerOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTThermometerOne0"><path fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M24 44a9 9 0 0 0 4-17.064V10c0-2 0-6-4-6s-4 4-4 6v16.936A9 9 0 0 0 24 44Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTThermometerOne0)"/>`),
		g.Group(children),
	)
}