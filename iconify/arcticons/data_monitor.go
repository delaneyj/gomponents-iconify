package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataMonitor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m4.5 33.4l9.155-11.57l8.282 16.035l6.693-27.73l4.234 8.168l6.146-6.821l4.49 8.405"/>`),
		g.Group(children),
	)
}