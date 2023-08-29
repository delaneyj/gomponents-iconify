package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarioLauncher(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="35.393" cy="15.893" r="8.107" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="12.607" cy="32.107" r="8.107" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.607 24a8.107 8.107 0 0 1 0-16.213h15.775a8.107 8.107 0 0 0 0 16.213h7.011a8.107 8.107 0 0 1 0 16.213H19.618a8.107 8.107 0 0 0 0-16.213h-7.011Z"/>`),
		g.Group(children),
	)
}