package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PortMicroUsbTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15.172 10.5a.497.497 0 0 1 .353.146l1.829 1.829a.499.499 0 0 1 .146.353V13a.5.5 0 0 1-.5.5H7a.5.5 0 0 1-.5-.5v-.172a.496.496 0 0 1 .146-.353l1.829-1.829a.497.497 0 0 1 .353-.146h6.344Zm0-1.5H8.828a2 2 0 0 0-1.414.586l-1.828 1.828A2 2 0 0 0 5 12.828V13a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2v-.172a2 2 0 0 0-.586-1.414l-1.828-1.828A2 2 0 0 0 15.172 9Z"/>`),
		g.Group(children),
	)
}