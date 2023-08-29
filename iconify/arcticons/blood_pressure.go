package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BloodPressure(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m34.602 28.571l8.93-10.03a2.83 2.83 0 0 0 .713-2.027l-.023-.428a2.828 2.828 0 0 0-.542-1.525l-3.549-4.847a2.828 2.828 0 0 0-.96-.83l-.17-.09a2.827 2.827 0 0 0-1.536-.32l-5.966.454a2.824 2.824 0 0 0-1.955 1.006l-5.204 6.223a.975.975 0 0 1-1.41.09l-6.68-6.165a2.908 2.908 0 0 0-1.246-.68l-1.445-.371a2.907 2.907 0 0 0-2.07.238l-7.042 3.675a1.294 1.294 0 0 0-.656 1.465l1.242 4.916c.148.584.446 1.12.866 1.552L24 39.534"/>`),
		g.Group(children),
	)
}