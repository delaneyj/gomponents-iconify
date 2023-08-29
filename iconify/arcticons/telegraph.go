package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Telegraph(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.5 24A21.5 21.5 0 1 1 24 45.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m12.73 23.465l18-8.014a1.173 1.173 0 0 1 1.638 1.243l-2.417 16.39a1.173 1.173 0 0 1-1.994.655l-2.982-3.007a21.122 21.122 0 0 0-10.102-5.674l-2.04-.486a.587.587 0 0 1-.104-1.107Z"/>`),
		g.Group(children),
	)
}