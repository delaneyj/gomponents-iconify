package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func McafeeSecurity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.02 34.976V5.5L24 13.024L39.98 5.5v29.476L24 42.5L8.02 34.976z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.546 30.71V15.827L24 20.278l9.454-4.451V30.71L24 35.161l-9.454-4.451z"/>`),
		g.Group(children),
	)
}