package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OfficeReader(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M44.029 16.259L24.52 4.367L4.378 16.26l20.143 11.89Zm-33.306 3.763l-6.768 4.17l20.566 11.775L44.03 24.192l-6.465-3.947"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M44.029 16.259L24.52 4.367L4.378 16.26l20.143 11.89Zm-33.306 3.763l-6.768 4.17l20.566 11.775L44.03 24.192l-6.465-3.947M11.012 28.08l-6.767 4.17L24.81 44.026L44.319 32.25l-6.465-3.947"/>`),
		g.Group(children),
	)
}