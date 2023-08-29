package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalAltTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.555 5.5c-8.641.015-15.634 7.033-15.618 15.674a15.647 15.647 0 0 0 2.427 8.343l-1.204 5.18a.729.729 0 0 0 .872.875l5.233-1.199c7.305 4.616 16.97 2.436 21.585-4.869c4.616-7.305 2.436-16.969-4.87-21.585A15.647 15.647 0 0 0 21.556 5.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.027 25.643a9.134 9.134 0 0 0 9.758 15.443l3.055.7c.306.07.58-.205.51-.51l-.704-3.025a9.134 9.134 0 0 0-7.701-14.02a9.135 9.135 0 0 0-4.918 1.412Z"/>`),
		g.Group(children),
	)
}