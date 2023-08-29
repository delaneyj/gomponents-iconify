package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElectricIron(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTElectricIron0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M5 18.712c0-.393.319-.712.713-.712h13.439C32.322 18 43 28.677 43 41.848v0a.152.152 0 0 1-.152.152H5V18.712ZM5 18V8h20"/><circle cx="15" cy="27" r="4" fill="#555"/><path stroke-linecap="round" d="M5 36h37"/><path stroke-linecap="round" stroke-linejoin="round" d="M29 20c1.333-2 5.4-6.4 9-4c3.5 2.334.5 8-1 10"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTElectricIron0)"/>`),
		g.Group(children),
	)
}