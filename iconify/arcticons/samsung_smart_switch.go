package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SamsungSmartSwitch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.035 23.802c-4.443-.91-7.785-4.84-7.785-9.552c0-5.385 4.365-9.75 9.75-9.75s9.75 4.365 9.75 9.75v3m-7.785 6.948c4.443.91 7.785 4.84 7.785 9.552c0 5.385-4.365 9.75-9.75 9.75s-9.75-4.365-9.75-9.75v-3"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.078 33.578L14.25 30.75l-2.828 2.828m19.5-19.156l2.828 2.828l2.828-2.828"/>`),
		g.Group(children),
	)
}