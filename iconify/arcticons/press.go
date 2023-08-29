package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Press(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.487 42.5v-37H19.245v37M42.5 5.5h-9.013"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.245 5.5h-2.929a10.816 10.816 0 0 0 0 21.632h2.93"/>`),
		g.Group(children),
	)
}