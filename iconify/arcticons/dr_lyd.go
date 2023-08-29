package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DrLyd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 6.904h4.055v25.579H5.5zm6.589 16.503h4.055v9.077h-4.055zm6.589-7.53h4.055v14.595h-4.055zm6.589 0h4.055v25.218h-4.055zm6.589 4.281h4.055v14.595h-4.055zm6.589-10.521H42.5v25.115h-4.055z"/>`),
		g.Group(children),
	)
}