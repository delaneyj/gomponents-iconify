package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElderLauncher(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 19.42a13.673 13.673 0 0 1-13.673 13.672h0A13.673 13.673 0 0 1 15.154 19.42h0A13.673 13.673 0 0 1 28.827 5.746h0A13.673 13.673 0 0 1 42.5 19.42Zm-6.405 22.834H6.047m22.741 0v-9.162M8.535 42.254V25.16"/><circle cx="8.535" cy="20.006" r="3.035" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.1 33.638v-6.982m4.87 6.982v-6.982m27.983-7.236A10.127 10.127 0 1 1 28.828 9.292a10.127 10.127 0 0 1 10.126 10.126Z"/>`),
		g.Group(children),
	)
}