package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerConnectionCableSplitCablesCableSplitDeviceComputerElectronicsCordsCordSplitter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="6.88" r="1.5"/><path d="M5.5 13.38h3m-3-2.5h3m.249-7.754L11.252.623l2 2l-2.502 2.504zM2.75.63l-2 2m2 2l2-2"/></g>`),
		g.Group(children),
	)
}