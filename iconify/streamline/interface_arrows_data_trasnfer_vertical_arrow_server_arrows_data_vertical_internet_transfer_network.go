package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsDataTrasnferVerticalArrowServerArrowsDataVerticalInternetTransferNetwork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m2.5 3.5l3-3v13m6-3l-3 3V.5"/>`),
		g.Group(children),
	)
}