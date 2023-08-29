package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Animezone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 7.11L4.5 40.89h39L24 7.11z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.55 24.59L24 16.71L12.47 36.68h9.33m12.4-11.89l-6.86 11.89h8.19l-4.1-7.1"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.09 24.59H19.45l.96 4.21h5.94l-6.98 12.09"/>`),
		g.Group(children),
	)
}