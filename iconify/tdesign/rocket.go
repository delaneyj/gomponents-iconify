package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rocket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.557 3.1H20.9v1.343a7 7 0 0 1-2.05 4.95l-7.557 7.556l-4.243-4.243l7.557-7.556a7 7 0 0 1 4.95-2.05Zm-1.554 9.968l2.26-2.261A9 9 0 0 0 22.9 4.443V1.1h-3.343a9 9 0 0 0-6.364 2.636l-2.261 2.26l-5.657-.707L1.04 9.524L14.475 22.96l4.235-4.235l-.707-5.656Zm-1.792 1.791l.393 3.143l-2.129 2.129l-1.768-1.768l3.504-3.504Zm-7.07-7.071l-3.505 3.504l-1.768-1.768l2.13-2.129l3.142.393Zm-3.505 9.16l-3.535 3.536L.687 19.07l3.535-3.535l1.414 1.414Zm2.829 2.83l-3.536 3.535l-1.414-1.414l3.535-3.536l1.415 1.414Z"/>`),
		g.Group(children),
	)
}